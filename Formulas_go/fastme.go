package main

// Fastme - Accurate and fast distance-based phylogeny inference program
// Homepage: http://www.atgc-montpellier.fr/fastme/

import (
	"fmt"
	
	"os/exec"
)

func installFastme() {
	// Método 1: Descargar y extraer .tar.gz
	fastme_tar_url := "https://gite.lirmm.fr/atgc/FastME/raw/v2.1.6.3/tarball/fastme-2.1.6.3.tar.gz"
	fastme_cmd_tar := exec.Command("curl", "-L", fastme_tar_url, "-o", "package.tar.gz")
	err := fastme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastme_zip_url := "https://gite.lirmm.fr/atgc/FastME/raw/v2.1.6.3/tarball/fastme-2.1.6.3.zip"
	fastme_cmd_zip := exec.Command("curl", "-L", fastme_zip_url, "-o", "package.zip")
	err = fastme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastme_bin_url := "https://gite.lirmm.fr/atgc/FastME/raw/v2.1.6.3/tarball/fastme-2.1.6.3.bin"
	fastme_cmd_bin := exec.Command("curl", "-L", fastme_bin_url, "-o", "binary.bin")
	err = fastme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastme_src_url := "https://gite.lirmm.fr/atgc/FastME/raw/v2.1.6.3/tarball/fastme-2.1.6.3.src.tar.gz"
	fastme_cmd_src := exec.Command("curl", "-L", fastme_src_url, "-o", "source.tar.gz")
	err = fastme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastme_cmd_direct := exec.Command("./binary")
	err = fastme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
