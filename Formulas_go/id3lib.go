package main

// Id3lib - ID3 tag manipulation
// Homepage: https://id3lib.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installId3lib() {
	// Método 1: Descargar y extraer .tar.gz
	id3lib_tar_url := "https://downloads.sourceforge.net/project/id3lib/id3lib/3.8.3/id3lib-3.8.3.tar.gz"
	id3lib_cmd_tar := exec.Command("curl", "-L", id3lib_tar_url, "-o", "package.tar.gz")
	err := id3lib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	id3lib_zip_url := "https://downloads.sourceforge.net/project/id3lib/id3lib/3.8.3/id3lib-3.8.3.zip"
	id3lib_cmd_zip := exec.Command("curl", "-L", id3lib_zip_url, "-o", "package.zip")
	err = id3lib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	id3lib_bin_url := "https://downloads.sourceforge.net/project/id3lib/id3lib/3.8.3/id3lib-3.8.3.bin"
	id3lib_cmd_bin := exec.Command("curl", "-L", id3lib_bin_url, "-o", "binary.bin")
	err = id3lib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	id3lib_src_url := "https://downloads.sourceforge.net/project/id3lib/id3lib/3.8.3/id3lib-3.8.3.src.tar.gz"
	id3lib_cmd_src := exec.Command("curl", "-L", id3lib_src_url, "-o", "source.tar.gz")
	err = id3lib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	id3lib_cmd_direct := exec.Command("./binary")
	err = id3lib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
