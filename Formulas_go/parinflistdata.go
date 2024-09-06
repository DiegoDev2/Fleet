package main

// PariNflistdata - Data files for nflist() in PARI/GP
// Homepage: https://pari.math.u-bordeaux.fr/packages.html

import (
	"fmt"
	
	"os/exec"
)

func installPariNflistdata() {
	// Método 1: Descargar y extraer .tar.gz
	parinflistdata_tar_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/nflistdata.tgz"
	parinflistdata_cmd_tar := exec.Command("curl", "-L", parinflistdata_tar_url, "-o", "package.tar.gz")
	err := parinflistdata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parinflistdata_zip_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/nflistdata.tgz"
	parinflistdata_cmd_zip := exec.Command("curl", "-L", parinflistdata_zip_url, "-o", "package.zip")
	err = parinflistdata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parinflistdata_bin_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/nflistdata.tgz"
	parinflistdata_cmd_bin := exec.Command("curl", "-L", parinflistdata_bin_url, "-o", "binary.bin")
	err = parinflistdata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parinflistdata_src_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/nflistdata.tgz"
	parinflistdata_cmd_src := exec.Command("curl", "-L", parinflistdata_src_url, "-o", "source.tar.gz")
	err = parinflistdata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parinflistdata_cmd_direct := exec.Command("./binary")
	err = parinflistdata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pari")
exec.Command("latte", "install", "pari")
}
