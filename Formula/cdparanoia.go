package main

// Cdparanoia - Audio extraction tool for sampling CDs
// Homepage: https://www.xiph.org/paranoia/

import (
	"fmt"
	
	"os/exec"
)

func installCdparanoia() {
	// Método 1: Descargar y extraer .tar.gz
	cdparanoia_tar_url := "https://downloads.xiph.org/releases/cdparanoia/cdparanoia-III-10.2.src.tgz"
	cdparanoia_cmd_tar := exec.Command("curl", "-L", cdparanoia_tar_url, "-o", "package.tar.gz")
	err := cdparanoia_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdparanoia_zip_url := "https://downloads.xiph.org/releases/cdparanoia/cdparanoia-III-10.2.src.tgz"
	cdparanoia_cmd_zip := exec.Command("curl", "-L", cdparanoia_zip_url, "-o", "package.zip")
	err = cdparanoia_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdparanoia_bin_url := "https://downloads.xiph.org/releases/cdparanoia/cdparanoia-III-10.2.src.tgz"
	cdparanoia_cmd_bin := exec.Command("curl", "-L", cdparanoia_bin_url, "-o", "binary.bin")
	err = cdparanoia_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdparanoia_src_url := "https://downloads.xiph.org/releases/cdparanoia/cdparanoia-III-10.2.src.tgz"
	cdparanoia_cmd_src := exec.Command("curl", "-L", cdparanoia_src_url, "-o", "source.tar.gz")
	err = cdparanoia_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdparanoia_cmd_direct := exec.Command("./binary")
	err = cdparanoia_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
