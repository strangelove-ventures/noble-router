cd proto
buf generate
cd ..

cp -r github.com/strangelove-ventures/noble-router/* ./
rm -rf github.com
