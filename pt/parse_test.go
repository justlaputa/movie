package pt

import "testing"

func TestParseHDC(t *testing.T) {
	titles := []string{
		"Operation.Red.Sea.2018.1080p.BluRay.x264.DTS-HD.MA.7.1-HDChina",     // normal 1080p
		"Big.Fish.and.Begonia.2016.BluRay.720p.x264.DTS-HDChina",             // normal 720p
		"Killer.Joe.2011.Blu-ray.Remux.1080p.AVC.DTS-HD.MA.5.1-OurBit",       // with Blu-ray
		"Metro.2013.720p.Blu-Ray.x264.DTS.HDCLUB",                            // without proper group
		"Mudbound.2017.1080p.NF.WEB-DL.DD5.1.x264-NTG",                       // web-dl
		"Made in Hong Kong 1997 720p BluRay x264-WiKi",                       // use space
		"A.One.and.a.Two.2000.720p.BluRay.x264.DTS-zzz@HDC",                  // with user@group
		"Another WolfCop 2017 BluRay Remux 1080p AVC DTS-HD MA 5.1[4.15 GB]", // without group
		"Man.in.Black.1997.UHDTV.4K.HEVC-HDCTV[7.33 GB]",                     // UHDTV
		"The.Longest.Nite.1998.HDTV.1080p.H264.AAC-luobo333[3.5 GB]",         // HDTV
	}

	expected := []MovieInfo{
		{"Operation Red Sea", 2018, "HDChina", Blueray, FHD, 0},
		{"Big Fish and Begonia", 2016, "HDChina", Blueray, HD, 0},
		{"Killer Joe", 2011, "OurBit", Blueray, FHD, 0},
		{"Metro", 2013, "", Blueray, HD, 0},
		{"Mudbound", 2017, "NTG", WebDL, FHD, 0},
		{"Made in Hong Kong", 1997, "WiKi", Blueray, HD, 0},
		{"A One and a Two", 2000, "HDC", Blueray, HD, 0},
		{"Another WolfCop", 2017, "", Blueray, FHD, 4150000000},
		{"Man in Black", 1997, "HDCTV", UHDTV, UHD4K, 7330000000},
		{"The Longest Nite", 1998, "luobo333", HDTV, FHD, 3500000000},
	}

	for i := range titles {
		result := ParseHDCTitle(titles[i])
		if result != expected[i] {
			t.Errorf("Faied on title %s\nExpected: %+v\nResult  : %+v", titles[i], expected[i], result)
		}
	}
}
