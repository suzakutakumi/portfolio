# /usr/bin/bash
if [ -e /etc/systemd/system/portfolio.service ]; then
    sudo rm -r /etc/systemd/system/portfolio.service
fi

if [ -e /opt/portfolio ]; then
    sudo rm -r /opt/portfolio/
fi

if [ -e /etc/systemd/system/portfolioRedirect.service ]; then
    sudo rm -r /etc/systemd/system/portfolioRedirect.service
fi

cd ./back&&go build main.go
cd ..

cd ./HTTP&&go build main.go
cd ..

sudo cp -r . /opt/portfolio
sudo cp ./portfolio.service /etc/systemd/system/
sudo cp ./HTTP/portfolioRedirect.service /etc/systemd/system/

sudo chmod 755 /opt/portfolio/back/main
sudo chmod 755 /opt/portfolio/HTTP/main

rm ./back/main
rm ./HTTP/main

sudo systemctl enable portfolio.service
sudo systemctl enable portfolioRedirect.service
