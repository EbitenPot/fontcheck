sunit init;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls;

type

  { TInit }

  TInit = class(TForm)
    Check: TButton;
    Group: TGroupBox;
    InitShow: TLabel;
    License: TMemo;
  private

  public

  end;

var
  Init: TInit;

implementation

{$R *.lfm}

end.

