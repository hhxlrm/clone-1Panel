import http from '@/api';
import { ReqPage, ResPage } from '../interface';
import { App } from '../interface/app';

export const SyncApp = () => {
    return http.post<any>('apps/sync', {});
};

export const SearchApp = (req: App.AppReq) => {
    return http.post<App.AppResPage>('apps/search', req);
};

export const GetApp = (id: number) => {
    return http.get<App.AppDTO>('apps/' + id);
};

export const GetAppDetail = (id: number, version: string) => {
    return http.get<App.AppDetail>(`apps/detail/${id}/${version}`);
};

export const InstallApp = (install: App.AppInstall) => {
    return http.post<any>('apps/install', install);
};

export const GetAppInstalled = (info: ReqPage) => {
    return http.post<ResPage<App.AppInstalled>>('apps/installed', info);
};

export const InstalledOp = (op: App.AppInstalledOp) => {
    return http.post<any>('apps/installed/op', op);
};

export const SyncInstalledApp = () => {
    return http.post<any>('apps/installed/sync', {});
};

export const GetAppService = (key: string | undefined) => {
    return http.get<any>(`apps/services/${key}`);
};

export const GetAppBackups = (info: App.AppBackupReq) => {
    return http.post<ResPage<App.AppBackup>>('apps/installed/backups', info);
};

export const DelAppBackups = (req: App.AppBackupDelReq) => {
    return http.post<any>('apps/installed/backups/del', req);
};

export const GetAppUpdateVersions = (id: number) => {
    return http.get<any>(`apps/installed/${id}/versions`);
};