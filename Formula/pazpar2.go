package main

// Pazpar2 - Metasearching middleware webservice
// Homepage: https://www.indexdata.com/resources/software/pazpar2/

import (
	"fmt"
	
	"os/exec"
)

func installPazpar2() {
	// Método 1: Descargar y extraer .tar.gz
	pazpar2_tar_url := "https://ftp.indexdata.com/pub/pazpar2/pazpar2-1.14.1.tar.gz"
	pazpar2_cmd_tar := exec.Command("curl", "-L", pazpar2_tar_url, "-o", "package.tar.gz")
	err := pazpar2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pazpar2_zip_url := "https://ftp.indexdata.com/pub/pazpar2/pazpar2-1.14.1.zip"
	pazpar2_cmd_zip := exec.Command("curl", "-L", pazpar2_zip_url, "-o", "package.zip")
	err = pazpar2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pazpar2_bin_url := "https://ftp.indexdata.com/pub/pazpar2/pazpar2-1.14.1.bin"
	pazpar2_cmd_bin := exec.Command("curl", "-L", pazpar2_bin_url, "-o", "binary.bin")
	err = pazpar2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pazpar2_src_url := "https://ftp.indexdata.com/pub/pazpar2/pazpar2-1.14.1.src.tar.gz"
	pazpar2_cmd_src := exec.Command("curl", "-L", pazpar2_src_url, "-o", "source.tar.gz")
	err = pazpar2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pazpar2_cmd_direct := exec.Command("./binary")
	err = pazpar2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: yaz")
	exec.Command("latte", "install", "yaz").Run()
}
