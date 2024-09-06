package main

// Gdbgui - Modern, browser-based frontend to gdb (gnu debugger)
// Homepage: https://www.gdbgui.com/

import (
	"fmt"
	
	"os/exec"
)

func installGdbgui() {
	// Método 1: Descargar y extraer .tar.gz
	gdbgui_tar_url := "https://files.pythonhosted.org/packages/f5/22/b26e8ee14c570768bfa85a7efe1a384c8b07fee7d966ee067bf9e8fa3033/gdbgui-0.15.2.0.tar.gz"
	gdbgui_cmd_tar := exec.Command("curl", "-L", gdbgui_tar_url, "-o", "package.tar.gz")
	err := gdbgui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdbgui_zip_url := "https://files.pythonhosted.org/packages/f5/22/b26e8ee14c570768bfa85a7efe1a384c8b07fee7d966ee067bf9e8fa3033/gdbgui-0.15.2.0.zip"
	gdbgui_cmd_zip := exec.Command("curl", "-L", gdbgui_zip_url, "-o", "package.zip")
	err = gdbgui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdbgui_bin_url := "https://files.pythonhosted.org/packages/f5/22/b26e8ee14c570768bfa85a7efe1a384c8b07fee7d966ee067bf9e8fa3033/gdbgui-0.15.2.0.bin"
	gdbgui_cmd_bin := exec.Command("curl", "-L", gdbgui_bin_url, "-o", "binary.bin")
	err = gdbgui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdbgui_src_url := "https://files.pythonhosted.org/packages/f5/22/b26e8ee14c570768bfa85a7efe1a384c8b07fee7d966ee067bf9e8fa3033/gdbgui-0.15.2.0.src.tar.gz"
	gdbgui_cmd_src := exec.Command("curl", "-L", gdbgui_src_url, "-o", "source.tar.gz")
	err = gdbgui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdbgui_cmd_direct := exec.Command("./binary")
	err = gdbgui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gdb")
exec.Command("latte", "install", "gdb")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
