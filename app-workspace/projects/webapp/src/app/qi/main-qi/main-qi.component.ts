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
  dataList = <any>[1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6];
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
          this.getData(0);
          this.getStaticCount();
        }
      }
    });
    this.getStaticCount();
    this.getData(0);
  }
  getStaticCount() {
    if (this.getObj.type == "quotation") {
      this.filter = this.qouationFilterFilter;
    } else {
      this.filter = this.invoiceFilterFilter;
    }
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
  getData(obj: any) {
    this.commonApiService.getDataQI(this.getObj).subscribe(res => {
      if (res.statusCOde == 200) {
        this.dataList = res.info;
        this.totalCount = res.count;
      }
    })
  }
}
