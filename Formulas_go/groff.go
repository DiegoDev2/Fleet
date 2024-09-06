package main

// Groff - GNU troff text-formatting system
// Homepage: https://www.gnu.org/software/groff/

import (
	"fmt"
	
	"os/exec"
)

func installGroff() {
	// Método 1: Descargar y extraer .tar.gz
	groff_tar_url := "https://ftp.gnu.org/gnu/groff/groff-1.23.0.tar.gz"
	groff_cmd_tar := exec.Command("curl", "-L", groff_tar_url, "-o", "package.tar.gz")
	err := groff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	groff_zip_url := "https://ftp.gnu.org/gnu/groff/groff-1.23.0.zip"
	groff_cmd_zip := exec.Command("curl", "-L", groff_zip_url, "-o", "package.zip")
	err = groff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	groff_bin_url := "https://ftp.gnu.org/gnu/groff/groff-1.23.0.bin"
	groff_cmd_bin := exec.Command("curl", "-L", groff_bin_url, "-o", "binary.bin")
	err = groff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	groff_src_url := "https://ftp.gnu.org/gnu/groff/groff-1.23.0.src.tar.gz"
	groff_cmd_src := exec.Command("curl", "-L", groff_src_url, "-o", "source.tar.gz")
	err = groff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	groff_cmd_direct := exec.Command("./binary")
	err = groff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: netpbm")
exec.Command("latte", "install", "netpbm")
	fmt.Println("Instalando dependencia: psutils")
exec.Command("latte", "install", "psutils")
	fmt.Println("Instalando dependencia: uchardet")
exec.Command("latte", "install", "uchardet")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
