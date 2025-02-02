package blob

import (
	"context"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

type Resolver struct {
}

func (*Resolver) Resolve(ctx context.Context, sourceResolver *v1alpha1.SourceResolver) (v1alpha1.ResolvedSourceConfig, error) {
	return v1alpha1.ResolvedSourceConfig{
		Blob: &v1alpha1.ResolvedBlobSource{
			URL:     sourceResolver.Spec.Source.Blob.URL,
			SubPath: sourceResolver.Spec.Source.SubPath,
		},
	}, nil
}

func (*Resolver) CanResolve(sourceResolver *v1alpha1.SourceResolver) bool {
	return sourceResolver.IsBlob()
}
