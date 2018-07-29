package pt

import "testing"

func TestParseHDC(t *testing.T) {
	titles := []string{
		"Operation.Red.Sea.2018.1080p.BluRay.x264.DTS-HD.MA.7.1-HDChina", // normal 1080p
		"Big.Fish.and.Begonia.2016.BluRay.720p.x264.DTS-HDChina",         // normal 720p
		"Killer.Joe.2011.Blu-ray.Remux.1080p.AVC.DTS-HD.MA.5.1-OurBit",   // with Blu-ray
		"Metro.2013.720p.Blu-Ray.x264.DTS.HDCLUB",                        // without proper group
		"Mudbound.2017.1080p.NF.WEB-DL.DD5.1.x264-NTG",                   // web-dl
		"Made in Hong Kong 1997 720p BluRay x264-WiKi",                   // use space
		"A.One.and.a.Two.2000.720p.BluRay.x264.DTS-zzz@HDC",              // with user@group
		"Another WolfCop 2017 BluRay Remux 1080p AVC DTS-HD MA 5.1",      // without group
		"Man.in.Black.1997.UHDTV.4K.HEVC-HDCTV",                          // UHDTV
		"The.Longest.Nite.1998.HDTV.1080p.H264.AAC-luobo333",             // HDTV
	}

	expected := []MovieInfo{
		{"Operation Red Sea", 2018, "HDChina", Blueray, FHD},
		{"Big Fish and Begonia", 2016, "HDChina", Blueray, HD},
		{"Killer Joe", 2011, "OurBit", Blueray, FHD},
		{"Metro", 2013, "", Blueray, HD},
		{"Mudbound", 2017, "NTG", WebDL, FHD},
		{"Made in Hong Kong", 1997, "WiKi", Blueray, HD},
		{"A One and a Two", 2000, "HDC", Blueray, HD},
		{"Another WolfCop", 2017, "", Blueray, FHD},
		{"Man in Black", 1997, "HDCTV", UHDTV, UHD4K},
		{"The Longest Nite", 1998, "luobo333", HDTV, FHD},
	}

	for i := range titles {
		result := ParseHDCTitle(titles[i])
		if result != expected[i] {
			t.Errorf("Faied on title %s\nExpected: %+v\nResult  : %+v", titles[i], expected[i], result)
		}
	}
}
