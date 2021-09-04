// {{page}}.ts
Page({
    data:{
        licNo: '',
        name: '',
        genderIndex: 0,
        genders: ['未知','男','女'],
        birthDate: '1990-01-01',
        licImgURL: undefined as string | undefined,
        state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED'
    },

    onUploadLic(){
        wx.chooseImage({
            success: res => {
                if(res.tempFilePaths.length >0) {
                    this.setData({
                        licImgURL: res.tempFilePaths[0]
                    })
                    //TODO: upload image
                    setTimeout(() => {
                        this.setData({
                            licNo: '3252452345',
                            name: '张三',
                            genderIndex: 1,
                            birthDate: '1989-12-02',
                        })
                    },1000)
                }
            },
        })
    },

    onGenderChange(e: any){
        this.setData({
            genderIndex: parseInt(e.detail.value)
        })
    },

    onBirthDateChange(e: any) {
        this.setData({
            birthDate: e.detail.value,
        })
    },

    onSubmit() {
        this.setData({
            state: 'PENDING'
        })
        setTimeout(() => {
            this.onLicVerified()
        },3000);
    },

    onResubmit() {
        this.setData({
            state: 'UNSUBMITTED',
            licImgURL: undefined,
        })
    },

    onLicVerified() {
        this.setData({
            state: 'VERIFIED',
        })
    }
})