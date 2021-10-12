# /usr/bin/bash
if [ -e /etc/systemd/system/portfolio.service ]; then
    sudo rm -r /etc/systemd/system/portfolio.service
fi

if [ -e /opt/portfolio ]; then
    sudo rm -r /opt/portfolio/
fi

go build main.go

sudo cp -r . /opt/portfolio/
sudo cp ./portfolio.service /etc/systemd/system/

sudo chmod 755 /opt/portfolio/main
rm ./main

sudo systemctl enable portfolio.service
