package main

// Libvoikko - Linguistic software and Finnish dictionary
// Homepage: https://voikko.puimula.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibvoikko() {
	// Método 1: Descargar y extraer .tar.gz
	libvoikko_tar_url := "https://www.puimula.org/voikko-sources/libvoikko/libvoikko-4.3.2.tar.gz"
	libvoikko_cmd_tar := exec.Command("curl", "-L", libvoikko_tar_url, "-o", "package.tar.gz")
	err := libvoikko_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvoikko_zip_url := "https://www.puimula.org/voikko-sources/libvoikko/libvoikko-4.3.2.zip"
	libvoikko_cmd_zip := exec.Command("curl", "-L", libvoikko_zip_url, "-o", "package.zip")
	err = libvoikko_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvoikko_bin_url := "https://www.puimula.org/voikko-sources/libvoikko/libvoikko-4.3.2.bin"
	libvoikko_cmd_bin := exec.Command("curl", "-L", libvoikko_bin_url, "-o", "binary.bin")
	err = libvoikko_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvoikko_src_url := "https://www.puimula.org/voikko-sources/libvoikko/libvoikko-4.3.2.src.tar.gz"
	libvoikko_cmd_src := exec.Command("curl", "-L", libvoikko_src_url, "-o", "source.tar.gz")
	err = libvoikko_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvoikko_cmd_direct := exec.Command("./binary")
	err = libvoikko_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: foma")
	exec.Command("latte", "install", "foma").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: hfstospell")
	exec.Command("latte", "install", "hfstospell").Run()
}
