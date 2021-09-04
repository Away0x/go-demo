// index.ts
// 获取应用实例

/*interface Marker {
    iconPath: string
    id: number
    latitude: number
    longitude: number
    width: number
    height: number
}

const defaultAvatar = '../../assets/images/car.png'*/
const initialLat = 29.761267625855936
const initialLng = 121.87264654736123

Page({

    isPageShowing: false,

    data: {
        avatarURL: '',
        setting: {
            skew: 0,
            rotate: 0,
            showLocation: true,
            showScale: true,
            subKey: '',
            layerStyle: 1,
            enableZoom: true,
            enableScroll: true,
            enableRotate: false,
            showCompass: false,
            enable3D: false,
            enableOverlooking: false,
            enableSatellite: false,
            enableTraffic: false,
        },
        location: {
            latitude: initialLat,
            longitude: initialLng,
        },
        scale: 10,
        markers: [
            {
              iconPath: '../../assets/images/car.png',
              id: 0,
              latitude: 23.09994,
              longitude: 113.324520,
              width: 50,
              height: 50
            },
            {
                iconPath: '../../assets/images/car.png',
                id: 1,
                latitude: 23.09994,
                longitude: 110.324520,
                width: 50,
                height: 50
            },
        ],
    },

   async onLoad() {
        const userInfo = await getApp<App.IAppOption>().globalData.userInfo
        this.setData({
            avatarURL: userInfo?.avatarUrl
        })
    },

    onMyLocationTap(){
        wx.getLocation({
            type: 'gcj02',
            success: res => {
                this.setData({
                    location: {
                        latitude: res.latitude,
                        longitude: res.longitude,
                    }
                })
            },
            fail: () => {
                wx.showToast({
                    icon: 'none',
                    title: '请前往设置页授权',
                })
            }
        })
    },

    onShow(): void | Promise<void> {
        this.isPageShowing = true
    },

    onHide(): void | Promise<void> {
        this.isPageShowing = false
    },

    moveCars() {
        const map = wx.createMapContext("map")
        const dest = {
            latitude: 23.09994,
            longitude: 110.324520,
        }
        const moveCar = () => {
            dest.latitude += 0.1
            dest.longitude += 0.1
            map.translateMarker({
                destination: {
                    latitude: dest.latitude,
                    longitude: dest.longitude,
                },
                markerId: 0,
                autoRotate: false,
                rotate: 0,
                duration: 5000,
                animationEnd:() => {
                    if(this.isPageShowing){
                        moveCar()
                    }
                }
            })
        }
        moveCar()
    },

    onScanClicked() {
        wx.scanCode({
            success:  ()=> {
                wx.navigateTo({
                    url: '/pages/register/register',
                })
            },
            fail : console.error,
        })
    },

    onMyTripsTap(){
        wx.navigateTo({
            url: '/pages/mytrips/mytrips',
        })
    }

})
