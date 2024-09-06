package main

// PariElldata - J.E. Cremona elliptic curve data for PARI/GP
// Homepage: https://pari.math.u-bordeaux.fr/packages.html

import (
	"fmt"
	
	"os/exec"
)

func installPariElldata() {
	// Método 1: Descargar y extraer .tar.gz
	parielldata_tar_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/elldata.tgz"
	parielldata_cmd_tar := exec.Command("curl", "-L", parielldata_tar_url, "-o", "package.tar.gz")
	err := parielldata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parielldata_zip_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/elldata.tgz"
	parielldata_cmd_zip := exec.Command("curl", "-L", parielldata_zip_url, "-o", "package.zip")
	err = parielldata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parielldata_bin_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/elldata.tgz"
	parielldata_cmd_bin := exec.Command("curl", "-L", parielldata_bin_url, "-o", "binary.bin")
	err = parielldata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parielldata_src_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/elldata.tgz"
	parielldata_cmd_src := exec.Command("curl", "-L", parielldata_src_url, "-o", "source.tar.gz")
	err = parielldata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parielldata_cmd_direct := exec.Command("./binary")
	err = parielldata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pari")
	exec.Command("latte", "install", "pari").Run()
}
