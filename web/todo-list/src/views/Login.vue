<template>
    <v-app id="inspire">
        <v-content>
            <v-container
                    class="fill-height"
                    fluid
            >
                <v-row
                        align="center"
                        justify="center"
                >
                    <v-col
                            cols="12"
                            sm="8"
                            md="4"
                    >
                        <v-card class="elevation-12">
                            <v-toolbar
                                    color="primary"
                                    dark
                                    flat
                            >
                                <v-toolbar-title>登录</v-toolbar-title>
                                <v-spacer />

                            </v-toolbar>
                            <v-card-text>
                                <v-form>
                                    <v-text-field
                                            v-model="loginForm.username"
                                            id="username"
                                            label="用户名"
                                            name="username"
                                            type="text"
                                    />

                                    <v-text-field
                                            v-model="loginForm.password"
                                            id="password"
                                            label="密码"
                                            name="password"
                                            type="password"
                                    />
                                </v-form>
                            </v-card-text>
                            <v-card-actions>
                                <v-spacer />
                                <v-btn color="primary" @click="login">登录</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
        </v-content>

        <v-snackbar
                v-model="message.status"
                :color=message.color
        >
            {{ message.text }}
            <v-btn
                    :color=message.color
                    @click="message.status = false"
            >
                关闭
            </v-btn>
        </v-snackbar>
    </v-app>
</template>

<script>
    export default {
        name: "Login",
        data(){
            return{
                loginForm:{
                    username:'',
                    password:'',
                },
                message:{
                    color:'success',
                    text:'',
                    status:false,
                    timeout:6000
                },
            }
        },
        methods:{
            login(){
                this.$api.login(this.loginForm)
                    .then(res=>{
                        if (res.code===20000){
                            var token = res.data.token;
                            localStorage.setItem("token",token)
                            this.message.text='新建任务成功';
                            this.message.status=true;
                            this.message.color='success';
                            this.$router.push({path: '/'})
                        }else {
                            this.message.text=res.message;
                            this.message.status=true;
                            this.message.color='error';
                        }

                    })
            }
        }
    }
</script>

<style scoped>

</style>