import { Pipe, PipeTransform } from '@angular/core';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import moment from 'moment';




@Pipe({ name: 'formatDateTime' })
export class FormatDateTime implements PipeTransform {

  transform(value: any) {
    return moment(value).format('lll');
  }
}


@Pipe({ name: 'dateToNowTime' })
export class DateToNowTime implements PipeTransform {

  transform(value: moment.MomentInput) {
    return moment.utc(value).utcOffset(new Date().getTimezoneOffset()).fromNow();
  }
}

@Pipe({ name: 'utcToLocalTime' })
export class UtcToLocalTime implements PipeTransform {

  transform(value: moment.MomentInput) {
    var stillUtc = moment.utc(value).toDate();
    return moment(stillUtc).local().format('lll');
  }
}

@Pipe({ name: 'utcToCurrentRelative' })
export class UtcToCurrentRelative implements PipeTransform {

  transform(value: any) {
    let t = moment.utc(value).utcOffset(value).toDate();
    return moment(moment(t).local()).calendar();
  }
}

@Pipe({ name: 'utcToCheckCurrentRelative' })
export class UtcToCheckCurrentRelative implements PipeTransform {

  transform(value: any) {
    let t = moment(moment.utc(value).utcOffset(value).toDate()).local();
    if (moment(t).isSame(moment(), 'day')) {
      return moment(t).format('LT');
    } else {
      return moment(t).format('L');
    }

  }
}

@Pipe({ name: 'safeHtml' })
export class Safe {
  constructor(private sanitizer: DomSanitizer) { }

  transform(style: string) {
    return this.sanitizer.bypassSecurityTrustHtml(style);
  }
}

@Pipe({ name: 'momentDatetime' })
export class MomentDate {
  constructor() { }

  transform(value: any) {
    return moment(value).format('lll');
  }
}

@Pipe({ name: 'notavailable' })
export class NotAvailable {
  constructor() { }
  transform(value: any) {
    return !!value ? value : 'Not available';
  }
}

@Pipe({ name: 'nacss' })
export class NotAvailableCSS {
  constructor() { }
  transform(value: any) {
    return !!value ? value : 'not-av';
  }
}

