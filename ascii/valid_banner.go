package fs

var validBanners = []string{"standard", "shadow", "thinkertoy"} // List of valid banner styles

func IsValidBanner(bannerStyle string) bool {
	for _, validBanner := range validBanners {
		if bannerStyle == validBanner {
			return true
		}
	}
	return false
}
