import { Component, OnInit, Output, EventEmitter, Input, } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { CommonApiService } from './../../services/common-api.service';
import moment from 'moment';

@Component({
  selector: 'app-create-qi',
  templateUrl: './create-qi.component.html',
  styleUrls: ['./create-qi.component.css']
})
export class CreateQiComponent implements OnInit {
  @Output() closeEmitter = new EventEmitter();
  @Input() type: any;
  @Input() inputInfo: any;
  @Input() taxable: any;
  dialogRef: any;
  body = <any>{
    date_t: new Date(),
    customer_name: '',
    customer_address: '',
    item_list: <any>[
    ],
    total: 0,
    terms_type: 0,
    terms_and_conditions: null
  };
  createObj: any;
  selectedIndex = -1;
  itemList = <any>[];
  customerList = <any>[];
  itemObj = { item_name: '', unit: 0, rate_per_item: 0, hsn_code: '', total: 0 }
  constructor(public dialog: MatDialog, private commonApiService: CommonApiService) { }

  ngOnInit(): void {
    if (!!this.inputInfo && !!this.inputInfo.id) {
      this.body = JSON.parse(JSON.stringify(this.inputInfo));
      this.body.date_t = new Date(this.inputInfo['date']);
    } else {
      this.body["tax_type"] = !!this.taxable ? 1 : 0;
    }
  }

  addItem(template: any) {
    this.body.item_list.push(JSON.parse(JSON.stringify(this.itemObj)));
    this.selectedIndex = this.body.item_list.length - 1;
  }

  calculateTotal(item: any) {
    if (!!item) {
      if (!!item.unit && !!item.rate_per_item) {
        item.total = item.unit * item.rate_per_item;
      } else {
        item.total = 0
      }
    }
    this.body.total = 0;
    this.body.item_list.forEach((i: any) => {
      if (!!i.unit && !!i.rate_per_item) {
        this.body.total = this.body.total + (i.unit * i.rate_per_item);
      }
    });
  }
  searchDialog(template: any, type: any) {
    this.dialogRef = this.dialog.open(template, {
      width: '350px',
    });
    if (type == 'customer') {
      this.searchCustomerName('');
      this.createObj = null;
      this.createObj = {
        customer_name: this.body.customer_name,
        address: this.body.customer_address,
      };

    } else {
      this.searchitemName('');
      let item = this.body.item_list[this.selectedIndex];
      this.createObj = {
        item_name: item['item_name'],
        hsn_code: item["hsn_code"],
        rate_per_item: item["rate_per_item"]
      };
    }
  }
  searchitemName(q: any) {
    if (!!q) {
      q = q.value;
    }
    let Obj = {
      limit: 20,
      q: q,
      type: 'rate_items'
    }
    this.commonApiService.getSearchOption(Obj).subscribe(res => {
      if (res.statusCode == 200) {
        this.itemList = res.info;
      }
    })
  }
  searchCustomerName(q: any) {
    if (!!q) {
      q = q.value;
    }
    let Obj = {
      limit: 15,
      q: q,
      type: 'customers'
    }
    this.commonApiService.getSearchOption(Obj).subscribe(res => {
      if (res.statusCode == 200) {
        this.customerList = res.info;
      }
    })
  }
  selectCusotmerInfo(item: any) {
    this.body.customer_name = item.customer_name;
    this.body.customer_address = item.address;
  }
  selectItemName(item: any) {
    if (this.body.item_list[this.selectedIndex].item_name == item.name) {
      setTimeout((i: any) => {
        this.body.item_list[this.selectedIndex].item_name = item.item_name;
      }, 50)

    } else {
      this.body.item_list[this.selectedIndex].item_name = item.item_name;
    }
    this.body.item_list[this.selectedIndex].rate_per_item = item.rate_per_item;
    this.body.item_list[this.selectedIndex].hsn_code = item.hsn_code;
    this.body.item_list[this.selectedIndex].unit = 1;
    this.body.item_list[this.selectedIndex].total = item.rate_per_item * item.unit;
    this.calculateTotal('');
  }
  create() {
    this.body['type'] = this.type;
    this.body["operation"] = !!this.inputInfo && !!this.inputInfo.id ? 'update' : 'create';
    this.body['date'] = moment(this.body.date_t).format('YYYY-MM-DD HH:mm:ss');
    if (!!this.inputInfo && !!this.inputInfo.id) {
      this.body['id'] = this.inputInfo.id
    }
    this.body.total = 0;
    this.body.item_list.forEach((i: any) => {
      if (!!i.unit && !!i.rate_per_item) {
        i.total = (i.unit * i.rate_per_item);
        this.body.total = this.body.total + (i.unit * i.rate_per_item);
      }
    });
    this.commonApiService.UpdateQI(this.body).subscribe(res => {
      this.closeEmitter.emit(true);
    })
  }



  deleteItem(index: number) {
    let item = this.body.item_list[index];
    if (!!item && !!item.id) {
      let table_name = '';
      if (this.type == 'invoice') {
        table_name = 'S_Invoice_Items'
      } else {
        table_name = 'S_Quotation_Items'
      }
      this.commonApiService.deleteRow(item.id, table_name).subscribe(res => {
        if (res.statusCode == 200) {
          this.body.item_list.splice(index, 1);
        }
      })
    } else {
      this.body.item_list.splice(index, 1);
    }

  }
  addItemMobile() {

  }

}
