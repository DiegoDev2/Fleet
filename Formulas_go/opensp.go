package main

// OpenSp - SGML parser
// Homepage: https://openjade.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installOpenSp() {
	// Método 1: Descargar y extraer .tar.gz
	opensp_tar_url := "https://downloads.sourceforge.net/project/openjade/opensp/1.5.2/OpenSP-1.5.2.tar.gz"
	opensp_cmd_tar := exec.Command("curl", "-L", opensp_tar_url, "-o", "package.tar.gz")
	err := opensp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opensp_zip_url := "https://downloads.sourceforge.net/project/openjade/opensp/1.5.2/OpenSP-1.5.2.zip"
	opensp_cmd_zip := exec.Command("curl", "-L", opensp_zip_url, "-o", "package.zip")
	err = opensp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opensp_bin_url := "https://downloads.sourceforge.net/project/openjade/opensp/1.5.2/OpenSP-1.5.2.bin"
	opensp_cmd_bin := exec.Command("curl", "-L", opensp_bin_url, "-o", "binary.bin")
	err = opensp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opensp_src_url := "https://downloads.sourceforge.net/project/openjade/opensp/1.5.2/OpenSP-1.5.2.src.tar.gz"
	opensp_cmd_src := exec.Command("curl", "-L", opensp_src_url, "-o", "source.tar.gz")
	err = opensp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opensp_cmd_direct := exec.Command("./binary")
	err = opensp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: docbook")
exec.Command("latte", "install", "docbook")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
