package main

// Zile - Text editor development kit
// Homepage: https://www.gnu.org/software/zile/

import (
	"fmt"
	
	"os/exec"
)

func installZile() {
	// Método 1: Descargar y extraer .tar.gz
	zile_tar_url := "https://ftp.gnu.org/gnu/zile/zile-2.6.2.tar.gz"
	zile_cmd_tar := exec.Command("curl", "-L", zile_tar_url, "-o", "package.tar.gz")
	err := zile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zile_zip_url := "https://ftp.gnu.org/gnu/zile/zile-2.6.2.zip"
	zile_cmd_zip := exec.Command("curl", "-L", zile_zip_url, "-o", "package.zip")
	err = zile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zile_bin_url := "https://ftp.gnu.org/gnu/zile/zile-2.6.2.bin"
	zile_cmd_bin := exec.Command("curl", "-L", zile_bin_url, "-o", "binary.bin")
	err = zile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zile_src_url := "https://ftp.gnu.org/gnu/zile/zile-2.6.2.src.tar.gz"
	zile_cmd_src := exec.Command("curl", "-L", zile_src_url, "-o", "source.tar.gz")
	err = zile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zile_cmd_direct := exec.Command("./binary")
	err = zile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libgee")
exec.Command("latte", "install", "libgee")
}
