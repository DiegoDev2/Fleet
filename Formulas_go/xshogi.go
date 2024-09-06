package main

// Xshogi - X11 interface for GNU Shogi
// Homepage: https://www.gnu.org/software/gnushogi/

import (
	"fmt"
	
	"os/exec"
)

func installXshogi() {
	// Método 1: Descargar y extraer .tar.gz
	xshogi_tar_url := "https://ftp.gnu.org/gnu/gnushogi/xshogi-1.4.2.tar.gz"
	xshogi_cmd_tar := exec.Command("curl", "-L", xshogi_tar_url, "-o", "package.tar.gz")
	err := xshogi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xshogi_zip_url := "https://ftp.gnu.org/gnu/gnushogi/xshogi-1.4.2.zip"
	xshogi_cmd_zip := exec.Command("curl", "-L", xshogi_zip_url, "-o", "package.zip")
	err = xshogi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xshogi_bin_url := "https://ftp.gnu.org/gnu/gnushogi/xshogi-1.4.2.bin"
	xshogi_cmd_bin := exec.Command("curl", "-L", xshogi_bin_url, "-o", "binary.bin")
	err = xshogi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xshogi_src_url := "https://ftp.gnu.org/gnu/gnushogi/xshogi-1.4.2.src.tar.gz"
	xshogi_cmd_src := exec.Command("curl", "-L", xshogi_src_url, "-o", "source.tar.gz")
	err = xshogi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xshogi_cmd_direct := exec.Command("./binary")
	err = xshogi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-shogi")
exec.Command("latte", "install", "gnu-shogi")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxaw")
exec.Command("latte", "install", "libxaw")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxmu")
exec.Command("latte", "install", "libxmu")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
}
