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
            cols="24"
            sm="6"
            md="6"
            lg="6"
          >
          <v-text-field
            label="任务"
            v-model=taskTitle
          ></v-text-field>
          
          </v-col>
          <v-col
            cols="24"
            sm="5"
            md="1"
            lg="1"
          >
            <v-btn  large color="cyan lighten-5" @click="createTask">添加任务</v-btn>
          </v-col>
          <v-col
            cols="24"
            sm="12"
            md="7"
            lg="7"
          >
                <v-card
                  width="100%"
                  class="mx-auto"
                >
                <v-toolbar
                  color="light-blue"
                  dark
                >
                  <v-toolbar-title>任务</v-toolbar-title>
                </v-toolbar>
                <v-list two-line subheader>
                  <v-list-item v-for="tasks in taskList" v-bind:key="tasks.id">
                  <template>
                    
                    <v-list-item-content>
                      <v-text-field
                        v-model="tasks.title"
                        :value="tasks.title"
                        v-if="tasks.status===true"
                        style="text-decoration:line-through"
                      ></v-text-field>
                        <v-text-field
                                v-model="tasks.title"
                                :value="tasks.title"
                                v-if="tasks.status===false"
                        ></v-text-field>
                    </v-list-item-content>
                    <v-btn small fab dark color="success" @click="updateStatus(tasks.id)"><v-icon dark>mdi-check</v-icon></v-btn>
                    <v-btn small fab dark color="primary" @click="updateTask(tasks.id,tasks.title)"><v-icon dark>mdi-update</v-icon></v-btn>
                    <v-btn small fab dark color="error" @click="saveDeleteTodo(tasks.id)"><v-icon dark>mdi-delete</v-icon></v-btn>
                  </template>
                    </v-list-item>
                        <v-list-item v-if="taskList.length===0">
                            <template>
                                <v-list-item-content>
                                    <v-text-field
                                            value="任务为空"
                                            disabled=true
                                    ></v-text-field>
                                </v-list-item-content>
                            </template>
                        </v-list-item>
                    </v-list>
              </v-card>
          </v-col>
        </v-row>

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

          <div>
              <v-row justify="center">
                  <v-dialog
                          v-model="dialog.status"
                          max-width="290"
                  >
                      <v-card>
                          <v-card-title class="headline">要删除这条任务吗？</v-card-title>

                          <v-card-actions>
                              <v-spacer></v-spacer>

                              <v-btn
                                      color="green darken-1"
                                      text
                                      @click="dialog.status = false"
                              >
                                  取消
                              </v-btn>

                              <v-btn
                                      color="error"
                                      text
                                      @click="deleteTodo(dialog.id),dialog.status = false"
                              >
                                  删除
                              </v-btn>
                          </v-card-actions>
                      </v-card>
                  </v-dialog>
              </v-row>
          </div>
      </v-container>
    </v-content>
  </v-app>
</template>


<script>
export default {
  name: 'Home',
  components: {
    
  },
  data(){
    return{
      taskTitle:'',
      taskList:[],
      message:{
          color:'success',
          text:'',
          status:false,
          timeout:6000
      },
      dialog: {
          status:false,
          id:0,
      }
    }
  },
  methods:{
    getAllTask(){
      this.$api.getAllTask()
        .then(res=>{
          this.taskList=res.data
        })
    },
    createTask(){
        var resource = {
            title:this.taskTitle
        };
        this.$api.createTodo(resource)
            .then(res=>{
                if (res.code===20000){
                    this.message.text='新建任务成功';
                    this.message.status=true;
                    this.message.color='success';
                    this.taskTitle='';
                    this.getAllTask()
                }else {
                    this.message.text=res.message;
                    this.message.status=true;
                    this.message.color='error';
                    this.getAllTask()
                }
            })
    },
    saveDeleteTodo(id){
      this.dialog.id=id
      this.dialog.status=true
    },
    deleteTodo(id){
        this.$api.deleteTodo(id)
            .then(res=>{
                if (res.code===20000){
                    this.message.text='删除任务成功';
                    this.message.status=true;
                    this.message.color='success';
                    this.getAllTask()
                }else {
                    this.message.text=res.message;
                    this.message.status=true;
                    this.message.color='error';
                }
            })

    },
    updateStatus(id){
        this.$api.updateStatus(id)
            .then(res=>{
                if (res.code===20000){
                    this.message.text='操作成功';
                    this.message.status=true;
                    this.message.color='success';
                    this.getAllTask()
                }else {
                    this.message.text=res.message;
                    this.message.status=true;
                    this.message.color='error';
                }
            })
    },
    updateTask(id,task){
        this.$api.UpdateTask(id,task)
            .then(res=>{
                if (res.code===20000){
                    this.message.text='更新成功';
                    this.message.status=true;
                    this.message.color='success';
                    this.getAllTask()
                }else {
                    this.message.text=res.message;
                    this.message.status=true;
                    this.message.color='error';
                }
            })
    }
  },
  mounted(){
    this.getAllTask()
  }
  
  
}
</script>
