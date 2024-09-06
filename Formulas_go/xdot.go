package main

// Xdot - Interactive viewer for graphs written in Graphviz's dot language
// Homepage: https://github.com/jrfonseca/xdot.py

import (
	"fmt"
	
	"os/exec"
)

func installXdot() {
	// Método 1: Descargar y extraer .tar.gz
	xdot_tar_url := "https://files.pythonhosted.org/packages/38/76/0503dddc3100e25135d1380f89cfa5d729b7d113a851804aa98dc4f19888/xdot-1.4.tar.gz"
	xdot_cmd_tar := exec.Command("curl", "-L", xdot_tar_url, "-o", "package.tar.gz")
	err := xdot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xdot_zip_url := "https://files.pythonhosted.org/packages/38/76/0503dddc3100e25135d1380f89cfa5d729b7d113a851804aa98dc4f19888/xdot-1.4.zip"
	xdot_cmd_zip := exec.Command("curl", "-L", xdot_zip_url, "-o", "package.zip")
	err = xdot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xdot_bin_url := "https://files.pythonhosted.org/packages/38/76/0503dddc3100e25135d1380f89cfa5d729b7d113a851804aa98dc4f19888/xdot-1.4.bin"
	xdot_cmd_bin := exec.Command("curl", "-L", xdot_bin_url, "-o", "binary.bin")
	err = xdot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xdot_src_url := "https://files.pythonhosted.org/packages/38/76/0503dddc3100e25135d1380f89cfa5d729b7d113a851804aa98dc4f19888/xdot-1.4.src.tar.gz"
	xdot_cmd_src := exec.Command("curl", "-L", xdot_src_url, "-o", "source.tar.gz")
	err = xdot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xdot_cmd_direct := exec.Command("./binary")
	err = xdot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: py3cairo")
exec.Command("latte", "install", "py3cairo")
	fmt.Println("Instalando dependencia: pygobject3")
exec.Command("latte", "install", "pygobject3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
