#!/bin/bash

set -e
echo 'export TERM=xterm-256color' >> ~/.bashrc

echo "Updating system..."
sudo apt update -y

echo "Installing Java 26..."
sudo apt install -y openjdk-26-jdk

echo "Verifying Java..."
java -version

echo "Creating server directory..."
mkdir -p ~/mc
cd ~/mc

echo "Downloading Minecraft server..."
curl -o server.jar https://piston-data.mojang.com/v1/objects/97ccd4c0ed3f81bbb7bfacddd1090b0c56f9bc51/server.jar

echo "Running server once to generate files..."
java -Xmx1G -Xms1G -jar server.jar nogui || true

echo "Accepting EULA..."
sed -i 's/eula=false/eula=true/' eula.txt

echo "Setting online-mode=true..."
sed -i 's/online-mode=.*/online-mode=true/' server.properties

echo "Creating start script..."
cat <<EOF > start.sh
#!/bin/bash
java -Xmx4G -Xms4G -jar server.jar nogui
EOF

chmod +x start.sh

echo "Setup complete. Run ./start.sh to start the server."
