<template>
    <div class="profile-page">
        <section class="section section-shaped section-lg my-0">
            <div class="shape shape-style-1 bg-gradient-green">
                <span></span>
                <span></span>
                <span></span>
                <span></span>
                <span></span>
                <span></span>
                <span></span>
                <span></span>
            </div>
        </section>
        <section class="section section-skew">
            <div class="container">
                <card shadow class="card-profile" no-body>
                    <div class="px-4">
                        <div class="mt-5 py-5 border-top">
                            <div class="row justify-content-center">
                                <div class="col-lg-12">
                                    <div class="tree3">
                                        <input class="tree-search-input mb-3" type="text" v-model="searchword" placeholder="search..."/>
                                        <button class=" tree-search-btn" type="button" @click="search">search</button>
                                        
                                        <v-tree ref='tree1' :canDeleteRoot="true" :data='treeData1' :draggable='true' :tpl='tpl' :halfcheck='true' :multiple="true" />
                                    </div>
                                </div>
                            </div>
                        </div>

                    </div>
                </card>
            </div>
        </section>
    </div>
</template>

<script>
import "vue-tree-halower/dist/halower-tree.min.css"; // you can customize the style of the tree
import { VTree, VSelectTree } from "vue-tree-halower";

export default {
  components: {
    VTree,
    VSelectTree
  },
  data() {
    return {
      searchword: "",
      initSelected: ["node-1"],
      treeData1: [
        {
          title: "node1",
          expanded: true,
          children: [
            {
              title: "node 1-1",
              expanded: true,
              children: [
                {
                  title: "node 1-1-1"
                },
                {
                  title: "node 1-1-2"
                },
                {
                  title: "node 1-1-3"
                }
              ]
            },
            {
              title: "node 1-2",
              children: [
                {
                  title: "<span style='color: red'>node 1-2-1</span>"
                },
                {
                  title: "<span style='color: red'>node 1-2-2</span>"
                }
              ]
            }
          ]
        }
      ],
      treeData2: [
        {
          title: "node1",
          expanded: false,
          async: true
        }
      ],

      treeData3: [
        {
          title: "node1",
          expanded: true,
          children: [
            {
              title: "node 1-1",
              expanded: true,
              children: [
                {
                  title: "node 1-1-1"
                },
                {
                  title: "node 1-1-2"
                },
                {
                  title: "node 1-1-3"
                }
              ]
            }
          ]
        }
      ]
    };
  },
  methods: {
    nodechecked(node, v) {
      alert("that a node-check envent ..." + node.title + v);
    },
    // tpl (node, ctx, parent, index, props) {
    tpl(...args) {
      let { 0: node, 2: parent, 3: index } = args;
      let titleClass = node.selected
        ? "node-title node-selected"
        : "node-title";
      if (node.searched) titleClass += " node-searched";
      return (
        <span>
          <button
            class="treebtn1"
            onClick={() =>
              this.$refs.tree1.addNode(node, {
                title: "sync node"
              })
            }
          >
            +
          </button>{" "}
          <span
            class={titleClass}
            domPropsInnerHTML={node.title}
            onClick={() => {
              this.$refs.tree1.nodeSelected(node);
            }}
          />{" "}
          <button class="treebtn2" onClick={() => this.asyncLoad1(node)}>
            async{" "}
          </button>{" "}
          <button
            class="treebtn3"
            onClick={() => this.$refs.tree1.delNode(node, parent, index)}
          >
            delete{" "}
          </button>{" "}
        </span>
      );
    },
    async asyncLoad1(node) {
      const { checked = false } = node;
      this.$set(node, "loading", true);
      let pro = new Promise(resolve => {
        setTimeout(resolve, 2000, ["async node1", "async node2"]);
      });
      this.$refs.tree1.addNodes(node, await pro);
      this.$set(node, "loading", false);
      if (checked) {
        this.$refs.tree1.childCheckedHandle(node, checked);
      }
    },
    async asyncLoad2(node) {
      const { checked = false } = node;
      this.$set(node, "loading", true);
      let pro = await new Promise(resolve => {
        setTimeout(resolve, 2000, [
          {
            title: "test1",
            async: true
          },
          {
            title: "test2",
            async: true
          },
          {
            title: "test3"
          }
        ]);
      });
      if (!node.hasOwnProperty("children")) {
        this.$set(node, "children", []);
      }
      node.children.push(...pro);
      this.$set(node, "loading", false);
      if (checked) {
        this.$refs.tree2.childCheckedHandle(node, checked);
      }
    },
    search() {
      this.$refs.tree1.searchNodes(this.searchword);
    }
  }
};
</script>

<style>
.treebtn1 {
  background-color: transparent;
  border: 1px solid #ccc;
  padding: 1px 3px;
  margin-right: 5px;
  color: rgb(148, 147, 147);
}
.treebtn2 {
  background-color: transparent;
  border: 1px solid #ccc;
  padding: 3px 5px;
  margin-left: 5px;
  color: rgb(97, 97, 97);
}
.treebtn3 {
  background-color: transparent;
  border: 1px solid #ccc;
  padding: 3px 5px;
  margin-left: 5px;
  color: rgb(63, 63, 63);
}
.tree-search-input {
  width: 70%;
  padding: 6px 8px;
  outline: none;
  border: 1px solid #ccc;
}
.tree-search-btn {
  width: 25%;
  padding: 6px 8px;
  outline: none;
  background-color: rgb(218, 218, 218);
  border: 1px solid rgb(226, 225, 225);
  color: rgb(117, 117, 117);
}
</style>
