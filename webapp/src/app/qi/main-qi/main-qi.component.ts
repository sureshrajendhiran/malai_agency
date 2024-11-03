import { Component } from '@angular/core';
import { DialogPosition, MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import {
  MatBottomSheet,
  MatBottomSheetModule,
  MatBottomSheetRef,
} from '@angular/material/bottom-sheet';
import { CommonApiService } from '../../services/common-api.service';

@Component({
  selector: 'app-main-qi',
  templateUrl: './main-qi.component.html',
  styleUrl: './main-qi.component.css'
})
export class MainQiComponent {
  filter = <any>[];
  qouationFilterFilter = <any>[
    { count: 0, name: "All", icon: "done_all" },
    { count: 0, name: "Sent", icon: "outgoing_mail" },
    { count: 0, name: "Pending", icon: "pending" },
    { count: 0, name: "Approved", icon: "check_box" },
    { count: 0, name: "Cancelled", icon: "block" }
  ];
  countObj: any;
  dataList = <any>[];
  totalCount: number = 100;
  isLoading: boolean = true;
  selectedItem: any;
  getObj = {
    type: <any>"quotation",
    limit: 20,
    page: 0,
    filter: <any>null,
    sort_type: null,
    static_type: "all"
  };
  type: any;
  invoiceFilterFilter = <any>[
    { count: 0, name: "All", icon: "done_all" },
    { count: 0, name: "Pending", icon: "pending" },
    { count: 0, name: "Sent", icon: "outgoing_mail" },
    { count: 0, name: "Sent and Pending", icon: "mark_email_unread" },
    { count: 0, name: "Paid", icon: "paid" },
    { count: 0, name: "Cancelled", icon: "block" }
  ];
  isLoadingPreview: boolean = false;
  previewData: any;
  routerSubcribe: any;
  dialogRef: any;
  bottomSheetDialogRef: any;
  taxable: number = 0;
  constructor(private route: Router,
    private _bottomSheet: MatBottomSheet,
    private dialog: MatDialog,
    private commonApiService: CommonApiService,
    private activeRoute: ActivatedRoute,) { }


  ngOnInit(): void {

    this.routerSubcribe = this.route.events.subscribe((event: any) => {
      if (event instanceof NavigationEnd) {
        let filterName = this.activeRoute.snapshot.paramMap.get('type');
        if (filterName != this.getObj.type) {
          this.getObj.type = filterName;
          this.getObj.page = 0;
          this.init();
        }
      }
    });
    this.getObj.type = this.activeRoute.snapshot.paramMap.get('type');
    this.init();

  }
  init() {
    if (this.getObj.type == "quotation") {
      this.filter = this.qouationFilterFilter;
    } else {
      this.filter = this.invoiceFilterFilter;
    }
    this.getObj.filter = this.filter[0];
    this.getStaticCount();
    this.getData(0);
  }
  getStaticCount() {
    this.commonApiService.getFilterCount(this.getObj.type, this.getObj.static_type).subscribe(res => {
      if (res.statusCode == 200) {
        this.filter.forEach((i: any) => {
          res.info.forEach((j: any) => {
            if (i.name == j.name) {
              i["count"] = !!j["count"] ? j["count"] : 0;
            }
          })
        });
      }
    });
  }

  OpenFilter(templateName: any) {
    const dialogPosition: DialogPosition = {
      top: '0px',
      right: '0px'
    };

    this.dialogRef = this.dialog.open(templateName, {
      width: '500px',
      height: '100%',
      position: dialogPosition
    });

    this.dialogRef.afterClosed().subscribe((res: any) => {
    });

  }

  openBottomSheet(tempName: any): void {
    this.bottomSheetDialogRef = this._bottomSheet.open(tempName,
    );
  }
  applyFilter(item: any) {
    if (!!this.bottomSheetDialogRef) {
      this.bottomSheetDialogRef.dismiss();
    }
    if (!!this.dialogRef) {
      this.dialogRef.close();
    }
    if (!!item) {
      this.getObj.filter = item;
      this.getObj.page = 0;
      this.getData(0);
    }
  }
  getData(page: number) {
    this.isLoading = true;
    this.getObj.page = page;
    this.commonApiService.getDataQI(this.getObj).subscribe(res => {
      if (res.statusCode == 200) {
        this.dataList = res.info;
        this.totalCount = res.count;
        this.isLoading = false;
      }
    })
  }
  changeType() {
    this.getData(0);
    this.getStaticCount();
  }
  handleOption(itemInfo: any, optionType: string) {

  }


  createDialog(templateName: any) {
    this.dialogRef = this.dialog.open(templateName, {
      width: '80%',
      height: '95%'
    });
  }
  handleOperation(action: any, templateName: any) {
    if (!!this.selectedItem) {
      if (action == 'edit_mobile') {
        this.openBottomSheet(templateName);
      } else
        if (action == "edit") {
          this.createDialog(templateName);
        } else if (action == 'preview' || action == 'preview_mobile' || action == 'download') {
          this.getQIInfoById(this.selectedItem.id, action, templateName)
        } else if (action == 'delete') {
          this.commonApiService.deleteRow(this.selectedItem.id, this.getObj.type == 'invoice' ? "S_Invoice" : "S_Quotation").subscribe(res => {
            if (res.statusCode == 200) {
              this.init();
            }
          })
        } else {
          let obj = {
            id: this.selectedItem.id,
            table_name: this.getObj.type == 'invoice' ? "S_Invoice" : "S_Quotation",
            status: action
          };
          this.commonApiService.updateCommon(obj, obj.table_name).subscribe(res => {
            if (res.statusCode == 200) {
              this.selectedItem['status'] = action;
              this.getStaticCount();
            }
          })
        }
    }
  }
  getQIInfoById(id: any, type: string, templateName: any) {
    this.isLoadingPreview = true;
    let tempType = type;
    if (type == "preview_mobile") {
      tempType = "preview"
    }
    this.commonApiService.getQIInfoById(this.getObj.type, tempType, id).subscribe(res => {
      if (res.statusCode == 200) {
        if (type == 'preview_mobile') {
          this.previewData = res.info;
          this.openBottomSheet(templateName);
        } else if (type == 'preview') {
          this.previewData = res.info;
        } else {
          this.commonApiService.savepdfFile(this.selectedItem.customer_name + "_" + this.getObj.type + "_" + this.selectedItem.ref_number, res.info);
        }
        this.isLoadingPreview = false;
      }
    });
  }


  close(e: any) {
    if (!!e) {
      this.init();
    }
    if (!!this.dialogRef) {
      this.dialogRef.close();
    }
    if (!!this._bottomSheet) {
      this._bottomSheet.dismiss();
    }
  }

}
