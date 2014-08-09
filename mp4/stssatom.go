//
// Copyright 2014 Hong Miao. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mp4

import (
	"log"
	//"errors"
	"github.com/oikomi/gomp4/util"
)

type StssAtom struct {
	Offset int64
	Size int64
	IsFullBox bool
	
	Version uint8
	Flag uint32
	EntriesNum uint32
	
	SyncSampleTable []uint32

	AllBytes []byte
}

func stssRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	var err error
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.Offset = offset
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.IsFullBox = false
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	size, _, err := fp.Mp4ReadHeader()
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	sizeInt := util.Bytes2Int(size)	
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.Size = sizeInt
		
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err := fp.Mp4Read(fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.Size)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.AllBytes = buf
		
	err = fp.Mp4Seek(offset + 8, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}	
			
	size, err = fp.Mp4Read(1)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.Version = uint8(size[0])
	
	size, err = fp.Mp4Read(3)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.Flag = util.Byte32Uint32(size, 0)	

	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.EntriesNum = util.Byte42Uint32(size, 0)
	
	var i uint32
	for i = 0; i < fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StssAtomAtomInstance.EntriesNum; i++ {	
		
		buf, err := fp.Mp4Read(4)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
			
		fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.StssAtomAtomInstance.SyncSampleTable = append(fs.MoovAtomInstance.
			TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.
			StssAtomAtomInstance.SyncSampleTable, util.Byte42Uint32(buf, 0))
	}
		
		
	return nil
	
}