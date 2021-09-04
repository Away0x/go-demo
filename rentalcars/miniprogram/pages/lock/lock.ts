/*import { IAppOption } from "../../appoption"
import { CarService } from "../../service/car"
import { car } from "../../service/proto_gen/car/car_pb"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"*/

const shareLocationKey = "share_location"

Page({
    carID: '',
    carRefresher: 0,

    data: {
        shareLocation: false,
        avatarURL: '',
    },

    async onLoad(){
        const userInfo = await getApp<App.IAppOption>().globalData.userInfo
        this.setData({
            avatarURL:userInfo?.avatarUrl,
            shareLocation: wx.getStorageSync(shareLocationKey),
        })
    },

    getUserProfile(){
      /*  wx.getUserProfile({
            desc:'用户获取用户头像',
            success:(res) => {
                this.setData({
                    avatarURL:res.userInfo.avatarUrl
                })
            }
        })*/
    },

    onShareLocation(e: any) {
        this.data.shareLocation = e.detail.value
        wx.setStorageSync(shareLocationKey, this.data.shareLocation)
    },

    onUnlockTap(){
        wx.getLocation({
            type: 'gcj02',
            success: loc => {
                console.log('starting a trip', {
                    location: {
                        latitude: loc.latitude,
                        longitude: loc.longitude,
                    },
                    avatarURL: this.data.shareLocation? this.data.avatarURL: ''
                })
            }
        })
        wx.showLoading({
            title: '开锁中',
            mask: true,
        })
        setTimeout(() => {
            wx.redirectTo({
                url: '/pages/driving/driving',
                complete: () => {
                    wx.hideLoading()
                }
            })
        }, 2000)
    },
    fail: () => {
        wx.showToast({
            icon: 'none',
            title: '请前往设置授权信息'
        })
    },
    /*onUnlockTap() {
        wx.getLocation({
            type: 'gcj02',
            success: async loc => {
                if (!this.carID) {
                    console.error('no carID specified')
                    return
                }
                let trip: rental.v1.ITripEntity
                try {
                    trip =  await TripService.createTrip({
                        start: loc,
                        carId: this.carID,
                        avatarUrl: this.data.shareLocation 
                                ? this.data.avatarURL : '',
                    })
                    if (!trip.id) {
                        console.error('no tripID in response', trip)
                        return
                    }
                } catch(err) {
                    wx.showToast({
                        title: '创建行程失败',
                        icon: 'none',
                    })
                    return
                }

                wx.showLoading({
                    title: '开锁中',
                    mask: true,
                })

                this.carRefresher = setInterval(async () => {
                    const c = await CarService.getCar(this.carID)
                    if (c.status === car.v1.CarStatus.UNLOCKED) {
                        this.clearCarRefresher()
                        wx.redirectTo({
                            url: routing.drving({
                                trip_id: trip.id!,
                            }),
                            complete: () => {
                                wx.hideLoading()
                            }
                        })
                    }
                }, 2000)
            },
            fail: () => {
                wx.showToast({
                    icon: 'none',
                    title: '请前往设置页授权位置信息',
                })
            }
        })
    },*/

    onUnload() {
        this.clearCarRefresher()
        wx.hideLoading()
    },

    clearCarRefresher() {
        if (this.carRefresher) {
            clearInterval(this.carRefresher)
            this.carRefresher = 0
        }
    },
})