package buddha

// TODO: These should probably all be private and have read only getters.
type Options struct {
	Seed int64
	Width int
	Height int
	
	PassCount int64
	MinIterations int
	MaxIterations int

	WorkerParrallelism int
	MergeParrallelism int
	
	RenderOptions *RenderOptions
	SaveOptions *SaveOptions
	LogOptions *LogOptions
}