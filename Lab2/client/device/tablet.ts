import Device from './device';
import requests from '../common/http';

export default class Tablet extends Device {
  constructor(name) {
    super(name);
  };
  public getData(count) {
    requests.getData(count);
  }
  private sendData = () => {
    requests.sendData(this.battery, this.currentVideo, new Date());
  }
  public startCycle() {
    const intervalVideo = setInterval(() => {
      this.currentVideo = Device.videos[Math.round(Math.random() * 4)];
    }, 8000);
    const intervalSend = setInterval(this.sendData, 10000);
    const intervalBattery = setInterval(() => {
      const battery = parseInt(this.battery);
      if (battery) {
        this.battery = (battery - 1).toString() + '%';
        return;
      }
      clearInterval(intervalBattery);
      clearInterval(intervalVideo);
      clearInterval(intervalSend);
    }, 2000);
  }
}
