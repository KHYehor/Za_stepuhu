export default abstract class Device {
  protected name: string = '';
  protected battery: string = '100%'
  protected currentVideo: string = '';
  constructor(name: string) {
    this.name = name;
  };
  static videos = ['video1', 'video2', 'video3', 'video4', 'video5'];
  abstract getData(count: number);
  abstract startCycle();
}

