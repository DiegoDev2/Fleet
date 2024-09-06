package main

// Zebra - Information management system
// Homepage: https://www.indexdata.com/resources/software/zebra/

import (
	"fmt"
	
	"os/exec"
)

func installZebra() {
	// Método 1: Descargar y extraer .tar.gz
	zebra_tar_url := "https://ftp.indexdata.com/pub/zebra/idzebra-2.2.7.tar.gz"
	zebra_cmd_tar := exec.Command("curl", "-L", zebra_tar_url, "-o", "package.tar.gz")
	err := zebra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zebra_zip_url := "https://ftp.indexdata.com/pub/zebra/idzebra-2.2.7.zip"
	zebra_cmd_zip := exec.Command("curl", "-L", zebra_zip_url, "-o", "package.zip")
	err = zebra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zebra_bin_url := "https://ftp.indexdata.com/pub/zebra/idzebra-2.2.7.bin"
	zebra_cmd_bin := exec.Command("curl", "-L", zebra_bin_url, "-o", "binary.bin")
	err = zebra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zebra_src_url := "https://ftp.indexdata.com/pub/zebra/idzebra-2.2.7.src.tar.gz"
	zebra_cmd_src := exec.Command("curl", "-L", zebra_src_url, "-o", "source.tar.gz")
	err = zebra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zebra_cmd_direct := exec.Command("./binary")
	err = zebra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: yaz")
exec.Command("latte", "install", "yaz")
}
