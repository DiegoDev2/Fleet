package main

// Srecord - Tools for manipulating EPROM load files
// Homepage: https://srecord.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSrecord() {
	// Método 1: Descargar y extraer .tar.gz
	srecord_tar_url := "https://downloads.sourceforge.net/project/srecord/srecord/1.64/srecord-1.64.tar.gz"
	srecord_cmd_tar := exec.Command("curl", "-L", srecord_tar_url, "-o", "package.tar.gz")
	err := srecord_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	srecord_zip_url := "https://downloads.sourceforge.net/project/srecord/srecord/1.64/srecord-1.64.zip"
	srecord_cmd_zip := exec.Command("curl", "-L", srecord_zip_url, "-o", "package.zip")
	err = srecord_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	srecord_bin_url := "https://downloads.sourceforge.net/project/srecord/srecord/1.64/srecord-1.64.bin"
	srecord_cmd_bin := exec.Command("curl", "-L", srecord_bin_url, "-o", "binary.bin")
	err = srecord_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	srecord_src_url := "https://downloads.sourceforge.net/project/srecord/srecord/1.64/srecord-1.64.src.tar.gz"
	srecord_cmd_src := exec.Command("curl", "-L", srecord_src_url, "-o", "source.tar.gz")
	err = srecord_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	srecord_cmd_direct := exec.Command("./binary")
	err = srecord_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: groff")
	exec.Command("latte", "install", "groff").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: groff")
	exec.Command("latte", "install", "groff").Run()
}
