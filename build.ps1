$GITBOOK_DIR="$PWD/gitbook"
$BUILD_DIR="$PWD/_dist"

gitbook install $GITBOOK_DIR
gitbook build $GITBOOK_DIR $BUILD_DIR