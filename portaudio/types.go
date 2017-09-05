// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 05 Sep 2017 19:49:49 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package portaudio

/*
#cgo pkg-config: portaudio-2.0
#include <portaudio.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Error type as declared in portaudio/portaudio.h:70
type Error int32

// DeviceIndex type as declared in portaudio/portaudio.h:161
type DeviceIndex int32

// HostApiIndex type as declared in portaudio/portaudio.h:187
type HostApiIndex int32

// HostApiInfo as declared in portaudio/portaudio.h:273
type HostApiInfo struct {
	StructVersion       int32
	Type                HostApiTypeId
	Name                string
	DeviceCount         int32
	DefaultInputDevice  DeviceIndex
	DefaultOutputDevice DeviceIndex
	ref2fd9b1fc         *C.PaHostApiInfo
	allocs2fd9b1fc      interface{}
}

// HostErrorInfo as declared in portaudio/portaudio.h:342
type HostErrorInfo struct {
	HostApiType    HostApiTypeId
	ErrorCode      int
	ErrorText      string
	ref5c7a77d7    *C.PaHostErrorInfo
	allocs5c7a77d7 interface{}
}

// Time type as declared in portaudio/portaudio.h:409
type Time float64

// SampleFormat type as declared in portaudio/portaudio.h:433
type SampleFormat uint

// DeviceInfo as declared in portaudio/portaudio.h:466
type DeviceInfo struct {
	StructVersion            int32
	Name                     string
	HostApi                  HostApiIndex
	MaxInputChannels         int32
	MaxOutputChannels        int32
	DefaultLowInputLatency   Time
	DefaultLowOutputLatency  Time
	DefaultHighInputLatency  Time
	DefaultHighOutputLatency Time
	DefaultSampleRate        float64
	refb60de337              *C.PaDeviceInfo
	allocsb60de337           interface{}
}

// StreamParameters as declared in portaudio/portaudio.h:530
type StreamParameters struct {
	Device                    DeviceIndex
	ChannelCount              int32
	SampleFormat              SampleFormat
	SuggestedLatency          Time
	HostApiSpecificStreamInfo unsafe.Pointer
	reff0e97a84               *C.PaStreamParameters
	allocsf0e97a84            interface{}
}

// Stream type as declared in portaudio/portaudio.h:584
type Stream [0]byte

// StreamFlags type as declared in portaudio/portaudio.h:602
type StreamFlags uint

// StreamCallbackTimeInfo as declared in portaudio/portaudio.h:652
type StreamCallbackTimeInfo struct {
	InputBufferAdcTime  Time
	CurrentTime         Time
	OutputBufferDacTime Time
	ref20d551f2         *C.PaStreamCallbackTimeInfo
	allocs20d551f2      interface{}
}

// StreamCallbackFlags type as declared in portaudio/portaudio.h:661
type StreamCallbackFlags uint

// StreamCallback type as declared in portaudio/portaudio.h:779
type StreamCallback func(input unsafe.Pointer, output unsafe.Pointer, frameCount uint, timeInfo *StreamCallbackTimeInfo, statusFlags StreamCallbackFlags, userData unsafe.Pointer) int32

// StreamFinishedCallback type as declared in portaudio/portaudio.h:911
type StreamFinishedCallback func(userData unsafe.Pointer)

// StreamInfo as declared in portaudio/portaudio.h:1018
type StreamInfo struct {
	StructVersion  int32
	InputLatency   Time
	OutputLatency  Time
	SampleRate     float64
	reff696a1d0    *C.PaStreamInfo
	allocsf696a1d0 interface{}
}
