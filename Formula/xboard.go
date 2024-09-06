package main

// Xboard - Graphical user interface for chess
// Homepage: https://www.gnu.org/software/xboard/

import (
	"fmt"
	
	"os/exec"
)

func installXboard() {
	// Método 1: Descargar y extraer .tar.gz
	xboard_tar_url := "https://ftp.gnu.org/gnu/xboard/xboard-4.9.1.tar.gz"
	xboard_cmd_tar := exec.Command("curl", "-L", xboard_tar_url, "-o", "package.tar.gz")
	err := xboard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xboard_zip_url := "https://ftp.gnu.org/gnu/xboard/xboard-4.9.1.zip"
	xboard_cmd_zip := exec.Command("curl", "-L", xboard_zip_url, "-o", "package.zip")
	err = xboard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xboard_bin_url := "https://ftp.gnu.org/gnu/xboard/xboard-4.9.1.bin"
	xboard_cmd_bin := exec.Command("curl", "-L", xboard_bin_url, "-o", "binary.bin")
	err = xboard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xboard_src_url := "https://ftp.gnu.org/gnu/xboard/xboard-4.9.1.src.tar.gz"
	xboard_cmd_src := exec.Command("curl", "-L", xboard_src_url, "-o", "source.tar.gz")
	err = xboard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xboard_cmd_direct := exec.Command("./binary")
	err = xboard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fairymax")
	exec.Command("latte", "install", "fairymax").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: librsvg")
	exec.Command("latte", "install", "librsvg").Run()
	fmt.Println("Instalando dependencia: polyglot")
	exec.Command("latte", "install", "polyglot").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}
