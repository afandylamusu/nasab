<template>
    <form role="form">
        <tabs fill class="flex-column flex-md-row">
            <card>
                <tab-pane>
                    <span slot="title">
                                                    <i class="ni ni-cloud-upload-96"></i>
                                                    Nasab
                                                </span>
                    <div class="description">
                        <div>
                            <h4 class="">Gender
                                <span></span>
                            </h4>
                            <span></span>
                            <base-radio name="male" class="mb-3" v-model="gender.male">Male</base-radio>
                            <base-radio name="female" v-model="gender.male">Female</base-radio>
                        </div>

                        <div>
                            <h4 class="mt-5">Names
                                <span></span>
                            </h4>
                            <span></span>
                            <base-input placeholder="first name"> </base-input>
                            <base-input placeholder="middle name"> </base-input>
                            <base-input placeholder="nick name"> </base-input>
                        </div>

                        <div>
                            <h4 class="mt-5">Family
                                <span></span>
                            </h4>
                            <base-input placeholder="kabilah"> </base-input>
                            <base-input placeholder="father"> </base-input>
                            <base-input placeholder="mother"> </base-input>
                        </div>

                    </div>
                </tab-pane>

                <tab-pane title="Profile">
                    <span slot="title">
                                                <i class="ni ni-bell-55 mr-2"></i>
                                                Profile
                                            </span>
                    <div class="description">

                        <h4 class="">Place, Date Of Birth</h4>
                        <div class="row">
                            <div class="col-lg-6 col-sm-6">
                                <base-input placeholder="place of birth">
                                </base-input>
                            </div>
                            <div class="col-lg-6 col-sm-6">
                                <base-input addon-left-icon="ni ni-calendar-grid-58">
                                    <flat-picker slot-scope="{focus, blur}" @on-open="focus" @on-close="blur" :config="{allowInput: true}" class="form-control datepicker" v-model="dates.simple">
                                    </flat-picker>
                                </base-input>
                            </div>
                        </div>

                        <h4 class="mt-5">Emails
                            <span class="text-right">&nbsp;<base-button outline type="success">Add<i class="ni ni-fat-add"></i></base-button></span>
                        </h4>
                        <base-input placeholder="email@example.com"> </base-input>

                        <h4 class="mt-5">Phones
                            <span class="text-right">&nbsp;<base-button outline type="success">Add<i class="ni ni-fat-add"></i></base-button></span>
                        </h4>
                        <base-input placeholder="+62 812 1234 1234"> </base-input>

                        <h4 class="mt-5">Homes
                            <span class="text-right">&nbsp;<base-button outline type="success">Add<i class="ni ni-fat-add"></i></base-button></span>
                        </h4>
                        <div class="row">
                            <div class="col-lg-5 col-sm-6">
                                <base-input placeholder="phone"> </base-input>
                            </div>
                            <div class="col-lg-7 col-sm-6">
                                <textarea class="form-control" id="exampleFormControlTextarea1" rows="3" placeholder="Write a large text here "></textarea>
                            </div>
                        </div>

                        <h4 class="mt-5">Offices
                            <span class="text-right">&nbsp;<base-button outline type="success">Add<i class="ni ni-fat-add"></i></base-button></span>
                        </h4>
                        <div class="row mt-4">
                            <div class="col-lg-5 col-sm-6">
                                <base-input placeholder="phone"> </base-input>
                            </div>
                            <div class="col-lg-7 col-sm-6">
                                <textarea class="form-control" id="exampleFormControlTextarea1" rows="3" placeholder="Write a large text here ..."></textarea>
                            </div>
                        </div>

                    </div>
                </tab-pane>

                <tab-pane>
                    <span slot="title">
                                                <i class="ni ni-calendar-grid-58"></i>
                                                Messages
                                            </span>
                    <p class="description">Raw denim you probably haven't heard of them jean shorts
                        Austin. Nesciunt tofu stumptown aliqua, retro synth master cleanse. Mustache
                        cliche tempor, williamsburg carles vegan helvetica. Reprehenderit butcher retro
                        keffiyeh dreamcatcher synth.</p>
                </tab-pane>
                <tab-pane>
                    <span slot="title">
                                                    <i class="ni ni-calendar-grid-58"></i>
                                                    Messages
                                                </span>
                    <p class="description">Raw denim you probably haven't heard of them jean shorts
                        Austin. Nesciunt tofu stumptown aliqua, retro synth master cleanse. Mustache
                        cliche tempor, williamsburg carles vegan helvetica. Reprehenderit butcher retro
                        keffiyeh dreamcatcher synth.</p>
                </tab-pane>
            </card>
        </tabs>

        <div class="mt-5">
            <base-button type="default">Save</base-button>
        </div>

    </form>
</template>

<script>
import flatPicker from "vue-flatpickr-component";
import "flatpickr/dist/flatpickr.css";
import Tabs from "@/components/Tabs/Tabs";
import TabPane from "@/components/Tabs/TabPane";

export default {
  components: {
    flatPicker,
    Tabs,
    TabPane
  },
  data() {
    return {
      dates: {
        simple: "2018-07-17"
      },
      gender: {
        male: "M",
        female: "F"
      }
    };
  }
};
</script>
