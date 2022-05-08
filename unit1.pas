unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ExtCtrls;

type

  { TForm1 }

  TForm1 = class(TForm)
    InFileButton: TButton;
    InFileDialog: TOpenDialog;
    KeyFileDialog: TOpenDialog;
    OutFileButton: TButton;
    KeyFileButton: TButton;
    EnButton: TCheckBox;
    InFile: TEdit;
    OutFile: TEdit;
    KeyFile: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    ComRadio: TRadioButton;
    UnComRadio: TRadioButton;
    RunButton: TToggleBox;
    procedure InFileButtonClick(Sender: TObject);
    procedure OutFileButtonClick(Sender: TObject);
    procedure KeyFileButtonClick(Sender: TObject);
    procedure EnButtonChange(Sender: TObject);
    procedure InFileChange(Sender: TObject);
    procedure OutFileChange(Sender: TObject);
    procedure KeyFileChange(Sender: TObject);
    procedure ComRadioChange(Sender: TObject);
    procedure UnComRadioChange(Sender: TObject);
    procedure RunButtonChange(Sender: TObject);
  private

  public

  end;

var
  Form1: TForm1;

implementation

{$R *.lfm}

{ TForm1 }

procedure TForm1.InFileButtonClick(Sender: TObject);
begin

end;

procedure TForm1.OutFileButtonClick(Sender: TObject);
begin

end;

procedure TForm1.KeyFileButtonClick(Sender: TObject);
begin

end;

procedure TForm1.EnButtonChange(Sender: TObject);
begin

end;

procedure TForm1.InFileChange(Sender: TObject);
begin

end;

procedure TForm1.OutFileChange(Sender: TObject);
begin

end;

procedure TForm1.KeyFileChange(Sender: TObject);
begin

end;

procedure TForm1.ComRadioChange(Sender: TObject);
begin

end;

procedure TForm1.UnComRadioChange(Sender: TObject);
begin

end;

procedure TForm1.RunButtonChange(Sender: TObject);
begin

end;

end.

