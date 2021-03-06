// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "通用授权模块",
    "title": "Passport(Api)",
    "version": "1.0.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/",
  "paths": {
    "/file/upload": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "FileUpload"
        ],
        "summary": "文件上传",
        "operationId": "/file/upload",
        "parameters": [
          {
            "type": "string",
            "description": "上传者的ID",
            "name": "euid",
            "in": "formData"
          },
          {
            "type": "file",
            "description": "上传的文件",
            "name": "file",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "成功，返回成功信息",
            "schema": {
              "properties": {
                "files": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/Photo"
                  }
                },
                "state": {
                  "$ref": "#/definitions/State"
                }
              }
            }
          },
          "default": {
            "description": "返回错误",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/passport/bindMobile": {
      "post": {
        "description": "绑定手机号接口",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Passport"
        ],
        "summary": "绑定手机号接口",
        "operationId": "/passport/bindMobile",
        "parameters": [
          {
            "description": "绑定手机号JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "phone",
                "valid_code"
              ],
              "properties": {
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "phone": {
                  "description": "手机号",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "valid_code": {
                  "description": "手机验证码",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "绑定成功，返回成功信息",
            "schema": {
              "$ref": "#/definitions/State"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/passport/getbackPwd": {
      "post": {
        "description": "找回密码接口",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Passport"
        ],
        "summary": "找回密码接口",
        "operationId": "/passport/getbackPwd",
        "parameters": [
          {
            "description": "找回密码JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "phone",
                "valid_code",
                "password"
              ],
              "properties": {
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "password": {
                  "description": "新密码",
                  "type": "string"
                },
                "phone": {
                  "description": "手机号",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "valid_code": {
                  "description": "手机验证码，快捷登录方式可用",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "返回成功信息",
            "schema": {
              "$ref": "#/definitions/State"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/passport/login": {
      "post": {
        "description": "登录接口",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Passport"
        ],
        "summary": "登录接口",
        "operationId": "/passport/login",
        "parameters": [
          {
            "description": "用户登录JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "user_name",
                "password"
              ],
              "properties": {
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "password": {
                  "description": "登录密码",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "user_name": {
                  "description": "登录账号",
                  "type": "string"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "登录成功，返回登录信息",
            "schema": {
              "$ref": "#/definitions/LoginState"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/passport/loginByThirdParty": {
      "post": {
        "description": "第三方登录",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Passport"
        ],
        "summary": "第三方登录",
        "operationId": "/passport/loginByThirdParty",
        "parameters": [
          {
            "description": "第三方登录JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "type",
                "open_id"
              ],
              "properties": {
                "avatar": {
                  "description": "第三方平台用户头像",
                  "type": "string"
                },
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "name": {
                  "description": "第三方平台用户用户名",
                  "type": "string"
                },
                "open_id": {
                  "description": "开放平台唯一ID",
                  "type": "string"
                },
                "platform": {
                  "description": "第三方登录平台，当type=3时，需要传入platform来判断从哪个平台登入",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "type": {
                  "description": "第三方登录方式 0=微信,1=QQ,2=微博,3=其他",
                  "type": "integer",
                  "format": "int64"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "返回登录成功信息",
            "schema": {
              "$ref": "#/definitions/LoginState"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/passport/quickLogin": {
      "post": {
        "description": "快捷登录接口",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Passport"
        ],
        "summary": "快捷登录接口",
        "operationId": "/passport/quickLogin",
        "parameters": [
          {
            "description": "快捷登录JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "phone",
                "valid_code"
              ],
              "properties": {
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "phone": {
                  "description": "手机号",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "valid_code": {
                  "description": "手机验证码，快捷登录方式可用",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "登录成功，返回登录信息",
            "schema": {
              "$ref": "#/definitions/LoginState"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/passport/register": {
      "post": {
        "description": "注册接口",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Passport"
        ],
        "summary": "注册接口",
        "operationId": "/passport/register",
        "parameters": [
          {
            "description": "用户注册JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "user_name",
                "password",
                "valid_code"
              ],
              "properties": {
                "birth_day": {
                  "description": "生日(格式必须为:yyyy-MM-dd)",
                  "type": "string"
                },
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "gender": {
                  "description": "用户性别 (0:保密 1:男 2:女)",
                  "type": "integer",
                  "format": "int64"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "invite_code": {
                  "description": "邀请码",
                  "type": "string"
                },
                "nick_name": {
                  "description": "用户昵称，最大11个字符",
                  "type": "string"
                },
                "password": {
                  "description": "登录密码",
                  "type": "string"
                },
                "platform": {
                  "description": "平台",
                  "type": "integer",
                  "format": "int64"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "user_name": {
                  "description": "注册账号",
                  "type": "string"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "valid_code": {
                  "description": "手机验证码，快捷登录方式可用",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "注册成功，返回注册成功信息",
            "schema": {
              "$ref": "#/definitions/RegisterState"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/passport/sendSms": {
      "post": {
        "description": "下发短息验证码",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Passport"
        ],
        "summary": "下发短信验证码",
        "operationId": "/passport/sendSms",
        "parameters": [
          {
            "description": "下发短信验证码请求JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "phone"
              ],
              "properties": {
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "phone": {
                  "description": "手机号",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "type": {
                  "description": "下发验证码的方式(0:注册 1:快捷登录 2:绑定手机号 3:找回密码)",
                  "type": "integer",
                  "format": "int64"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "验证码发送成功",
            "schema": {
              "$ref": "#/definitions/State"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/getOwnerAccount": {
      "get": {
        "tags": [
          "User"
        ],
        "summary": "查看用户账户信息",
        "operationId": "/user/getOwnerAccount",
        "parameters": [
          {
            "type": "string",
            "description": "加密后的用户ID",
            "name": "euid",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "返回成功信息",
            "schema": {
              "$ref": "#/definitions/AccountState"
            }
          },
          "default": {
            "description": "返回错误",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/updateAvatar": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "summary": "修改头像API",
        "operationId": "/user/updateAvatar",
        "parameters": [
          {
            "description": "修改头像JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "euid",
                "avatar"
              ],
              "properties": {
                "avatar": {
                  "description": "上传的头像图片地址",
                  "type": "string"
                },
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "euid": {
                  "description": "创建人ID",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "修改成功，返回成功信息",
            "schema": {
              "properties": {
                "state": {
                  "$ref": "#/definitions/AvatarState"
                }
              }
            }
          },
          "default": {
            "description": "返回错误",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/updateProfile": {
      "post": {
        "description": "修改用户资料接口",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "summary": "修改用户资料接口",
        "operationId": "/user/updateProfile",
        "parameters": [
          {
            "description": "修改用户资料JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "euid"
              ],
              "properties": {
                "birth_day": {
                  "description": "生日(格式必须为:yyyy-MM-dd)",
                  "type": "string"
                },
                "blood": {
                  "description": "血型",
                  "type": "string"
                },
                "company_name": {
                  "description": "公司名称",
                  "type": "string"
                },
                "degree": {
                  "description": "最高学历",
                  "type": "string"
                },
                "euid": {
                  "description": "用户ID",
                  "type": "string"
                },
                "gender": {
                  "description": "用户性别(0:保密 1:男 2:女)",
                  "type": "integer",
                  "format": "int64"
                },
                "home_area": {
                  "description": "家乡",
                  "type": "string"
                },
                "interest": {
                  "description": "兴趣爱好",
                  "type": "string"
                },
                "marriage": {
                  "description": "婚姻状况",
                  "type": "string"
                },
                "nick_name": {
                  "description": "用户昵称",
                  "type": "string"
                },
                "now_area": {
                  "description": "现居城市",
                  "type": "string"
                },
                "profession": {
                  "description": "职业信息",
                  "type": "string"
                },
                "remark": {
                  "description": "个人说明",
                  "type": "string"
                },
                "salary": {
                  "description": "目前月薪",
                  "type": "string"
                },
                "school": {
                  "description": "毕业学校",
                  "type": "string"
                },
                "shape": {
                  "description": "体型",
                  "type": "string"
                },
                "stature": {
                  "description": "身高",
                  "type": "integer",
                  "format": "int64"
                },
                "weight": {
                  "description": "体重",
                  "type": "integer",
                  "format": "int64"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "修改成功，返回成功信息",
            "schema": {
              "$ref": "#/definitions/UserState"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/updatePwd": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "summary": "修改密码接口",
        "operationId": "/user/updatePwd",
        "parameters": [
          {
            "description": "修改密码JSON",
            "name": "body",
            "in": "body",
            "schema": {
              "required": [
                "euid",
                "old_pwd",
                "new_pwd"
              ],
              "properties": {
                "client": {
                  "description": "客户端类型",
                  "type": "string"
                },
                "euid": {
                  "description": "用户ID",
                  "type": "string"
                },
                "imei": {
                  "description": "唯一识别号",
                  "type": "string"
                },
                "new_pwd": {
                  "description": "新密码",
                  "type": "string"
                },
                "old_pwd": {
                  "description": "旧密码",
                  "type": "string"
                },
                "phone": {
                  "description": "手机号",
                  "type": "string"
                },
                "ts": {
                  "description": "时间戳(加密串要用到,供服务端验证，简单防刷)",
                  "type": "integer",
                  "format": "int64"
                },
                "val": {
                  "description": "加密串",
                  "type": "string"
                },
                "version": {
                  "description": "版本号",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "修改成功，返回成功信息",
            "schema": {
              "$ref": "#/definitions/State"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Account": {
      "description": "返回的账户信息",
      "type": "object",
      "properties": {
        "current_coins": {
          "description": "用户当前金币",
          "type": "integer",
          "format": "int64"
        },
        "current_points": {
          "description": "用户当前积分",
          "type": "integer",
          "format": "int64"
        },
        "current_rmb": {
          "description": "用户当前人民币",
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "AccountState": {
      "description": "获取用户账户信息成功后的信息",
      "type": "object",
      "required": [
        "state"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/Account"
        },
        "state": {
          "$ref": "#/definitions/State"
        }
      }
    },
    "AvatarState": {
      "description": "上传头像成功后的信息",
      "type": "object",
      "required": [
        "state"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/Photo"
        },
        "state": {
          "$ref": "#/definitions/State"
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "msg"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "msg": {
          "description": "返回结果",
          "type": "string"
        }
      }
    },
    "LoginState": {
      "description": "登录成功后返回的信息",
      "type": "object",
      "required": [
        "state"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/UserBase"
        },
        "state": {
          "$ref": "#/definitions/State"
        }
      }
    },
    "Photo": {
      "description": "返回的图片信息",
      "type": "object",
      "required": [
        "photo_id"
      ],
      "properties": {
        "large": {
          "description": "大图",
          "type": "string"
        },
        "middle": {
          "description": "中图",
          "type": "string"
        },
        "origin": {
          "description": "原图",
          "type": "string"
        },
        "photo_id": {
          "description": "图片ID",
          "type": "integer",
          "format": "int64"
        },
        "round": {
          "description": "圆图(头像)",
          "type": "string"
        },
        "small": {
          "description": "小图",
          "type": "string"
        },
        "type": {
          "description": "文件类型，后缀",
          "type": "string"
        }
      }
    },
    "RegisterState": {
      "description": "注册成功后返回的信息",
      "type": "object",
      "required": [
        "state"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/UserBase"
        },
        "state": {
          "$ref": "#/definitions/State"
        }
      }
    },
    "State": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserBase": {
      "description": "登录成功返回的信息",
      "type": "object",
      "required": [
        "euid",
        "nick_name"
      ],
      "properties": {
        "avatar": {
          "description": "用户头像",
          "type": "string"
        },
        "birth_day": {
          "description": "生日(格式为:yyyy-MM-dd)",
          "type": "string"
        },
        "current_coins": {
          "description": "用户当前金币",
          "type": "integer",
          "format": "int64"
        },
        "current_points": {
          "description": "用户当前积分",
          "type": "integer",
          "format": "int64"
        },
        "euid": {
          "description": "用户ID",
          "type": "string"
        },
        "friend_count": {
          "description": "好友数量",
          "type": "integer",
          "format": "int64"
        },
        "gender": {
          "description": "性别",
          "type": "integer",
          "format": "int64"
        },
        "level": {
          "description": "用户在应用内的等级",
          "type": "integer",
          "format": "int64"
        },
        "login_at": {
          "description": "上一次登录日期",
          "type": "integer",
          "format": "int64"
        },
        "nick_name": {
          "description": "用户昵称",
          "type": "string"
        },
        "register_at": {
          "description": "注册日期",
          "type": "integer",
          "format": "int64"
        },
        "tags": {
          "description": "标签",
          "type": "string"
        }
      }
    },
    "UserInfo": {
      "description": "返回的用户信息",
      "type": "object",
      "required": [
        "euid",
        "nick_name"
      ],
      "properties": {
        "avatar": {
          "description": "用户头像",
          "type": "string"
        },
        "birth_day": {
          "description": "生日(格式为:yyyy-MM-dd)",
          "type": "string"
        },
        "blood": {
          "description": "血型",
          "type": "string"
        },
        "company_name": {
          "description": "公司名称",
          "type": "string"
        },
        "degree": {
          "description": "最高学历",
          "type": "string"
        },
        "euid": {
          "description": "用户ID",
          "type": "string"
        },
        "gender": {
          "description": "性别",
          "type": "integer",
          "format": "int64"
        },
        "home_area": {
          "description": "家乡",
          "type": "string"
        },
        "interest": {
          "description": "兴趣爱好",
          "type": "string"
        },
        "marriage": {
          "description": "婚姻状况",
          "type": "string"
        },
        "nick_name": {
          "description": "用户昵称",
          "type": "string"
        },
        "now_area": {
          "description": "现居城市",
          "type": "string"
        },
        "profession": {
          "description": "职业信息",
          "type": "string"
        },
        "resume": {
          "description": "个人简介",
          "type": "string"
        },
        "salary": {
          "description": "目前月薪",
          "type": "string"
        },
        "school": {
          "description": "毕业学校",
          "type": "string"
        },
        "shape": {
          "description": "体型",
          "type": "string"
        },
        "stature": {
          "description": "身高(cm)",
          "type": "string"
        },
        "tags": {
          "description": "标签",
          "type": "string"
        },
        "weight": {
          "description": "体重(kg)",
          "type": "string"
        }
      }
    },
    "UserState": {
      "description": "修改用户信息成功后的信息",
      "type": "object",
      "required": [
        "state"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/UserInfo"
        },
        "state": {
          "$ref": "#/definitions/State"
        }
      }
    }
  },
  "tags": [
    {
      "description": "用户授权",
      "name": "Passport"
    },
    {
      "description": "用户信息",
      "name": "User"
    }
  ]
}`))
}
