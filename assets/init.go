package assets

func init() {

	// regist grandet's assets
	(&grandetsAssets{}).registAssets()

	// initialize templates
	asset_go_tmpl = getTemplate("asset.go.tmpl")
	grandet_go_tmpl = getTemplate("grandet.go.tmpl")
}
