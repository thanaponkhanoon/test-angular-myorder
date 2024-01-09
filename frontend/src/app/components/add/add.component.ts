import {Component} from '@angular/core';
import {MatTableModule} from '@angular/material/table';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatIconModule} from '@angular/material/icon';
import {MatDividerModule} from '@angular/material/divider';
import {MatButtonModule} from '@angular/material/button';
import {
  MatDialog,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogTitle,
} from '@angular/material/dialog';
import {MatSelectModule} from '@angular/material/select';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatToolbarModule} from '@angular/material/toolbar';

export interface PeriodicElement {
  name: string;
  id: number;
  u: string;
  c: string;
}

const ELEMENT_DATA: PeriodicElement[] = [
  {id: 1, name: 'ปากกา', u: 'แท่ง', c: 'เครื่องเขียน'},
  {id: 2, name: 'โต๊ะ', u: 'ตัว', c: 'เฟอนิเจอร์'},
  {id: 3, name: 'โน๊ตบุ๊ค', u: 'เครื่อง', c: 'อุปกรณ์อิเล็กทรอนิกส์'},
];

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css'],
  standalone: true,
  imports: [MatTableModule, MatGridListModule,MatButtonModule, MatDividerModule, MatIconModule, MatToolbarModule],
})
export class AddComponent {
  displayedColumns: string[] = ['id', 'name', 'u', 'c', 'actions'];
  dataSource = ELEMENT_DATA;
  constructor(public dialog: MatDialog) {}

  openDialog() {
    this.dialog.open(DialogElementsExampleDialog);
  }
}

@Component({
  selector: 'add',
  templateUrl: 'edit.html',
  standalone: true,
  imports: [MatDialogTitle, MatDialogContent, MatDialogActions, MatDialogClose, MatButtonModule, 
    MatFormFieldModule, MatInputModule, MatSelectModule, MatGridListModule],
})
export class DialogElementsExampleDialog {}