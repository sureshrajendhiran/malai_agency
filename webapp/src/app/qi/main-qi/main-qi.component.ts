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
  getObj = {
    type: <any>"quotation",
    limit: 20,
    page: 0,
    filter: <any>null,
    sort_type: null,
    static_type: "all"
  };
  invoiceFilterFilter = <any>[
    { count: 0, name: "All", icon: "done_all" },
    { count: 0, name: "Pending", icon: "pending" },
    { count: 0, name: "Sent", icon: "outgoing_mail" },
    { count: 0, name: "Sent and Pending", icon: "mark_email_unread" },
    { count: 0, name: "Paid", icon: "paid" },
    { count: 0, name: "Cancelled", icon: "block" }
  ];
  previewData: any;
  routerSubcribe: any;
  dialogRef: any;
  bottomSheetDialogRef: any;
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
}
