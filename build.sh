echo "start build"
fyne package -icon ./res-static/icon/icon.png

echo "copy files to build-target"
mkdir ./build-target

echo "move build target..."
mv ./WirelessAdbConnect build-target/
echo ""

echo "copy resource"
cp -rf ./res-static/ build-target/
echo ""

echo "build success, you can open binary file in build-target"