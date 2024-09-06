package main

// QtUnixodbc - Qt SQL Database Driver
// Homepage: https://www.qt.io/

import (
	"fmt"
	
	"os/exec"
)

func installQtUnixodbc() {
	// Método 1: Descargar y extraer .tar.gz
	qtunixodbc_tar_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtunixodbc_cmd_tar := exec.Command("curl", "-L", qtunixodbc_tar_url, "-o", "package.tar.gz")
	err := qtunixodbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qtunixodbc_zip_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtunixodbc_cmd_zip := exec.Command("curl", "-L", qtunixodbc_zip_url, "-o", "package.zip")
	err = qtunixodbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qtunixodbc_bin_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtunixodbc_cmd_bin := exec.Command("curl", "-L", qtunixodbc_bin_url, "-o", "binary.bin")
	err = qtunixodbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qtunixodbc_src_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtunixodbc_cmd_src := exec.Command("curl", "-L", qtunixodbc_src_url, "-o", "source.tar.gz")
	err = qtunixodbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qtunixodbc_cmd_direct := exec.Command("./binary")
	err = qtunixodbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
}
