package main

// Wxpython - Python bindings for wxWidgets
// Homepage: https://www.wxpython.org/

import (
	"fmt"
	
	"os/exec"
)

func installWxpython() {
	// Método 1: Descargar y extraer .tar.gz
	wxpython_tar_url := "https://files.pythonhosted.org/packages/aa/64/d749e767a8ce7bdc3d533334e03bb1106fc4e4803d16f931fada9007ee13/wxPython-4.2.1.tar.gz"
	wxpython_cmd_tar := exec.Command("curl", "-L", wxpython_tar_url, "-o", "package.tar.gz")
	err := wxpython_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wxpython_zip_url := "https://files.pythonhosted.org/packages/aa/64/d749e767a8ce7bdc3d533334e03bb1106fc4e4803d16f931fada9007ee13/wxPython-4.2.1.zip"
	wxpython_cmd_zip := exec.Command("curl", "-L", wxpython_zip_url, "-o", "package.zip")
	err = wxpython_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wxpython_bin_url := "https://files.pythonhosted.org/packages/aa/64/d749e767a8ce7bdc3d533334e03bb1106fc4e4803d16f931fada9007ee13/wxPython-4.2.1.bin"
	wxpython_cmd_bin := exec.Command("curl", "-L", wxpython_bin_url, "-o", "binary.bin")
	err = wxpython_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wxpython_src_url := "https://files.pythonhosted.org/packages/aa/64/d749e767a8ce7bdc3d533334e03bb1106fc4e4803d16f931fada9007ee13/wxPython-4.2.1.src.tar.gz"
	wxpython_cmd_src := exec.Command("curl", "-L", wxpython_src_url, "-o", "source.tar.gz")
	err = wxpython_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wxpython_cmd_direct := exec.Command("./binary")
	err = wxpython_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: sip")
exec.Command("latte", "install", "sip")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: six")
exec.Command("latte", "install", "six")
	fmt.Println("Instalando dependencia: wxwidgets")
exec.Command("latte", "install", "wxwidgets")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
}
