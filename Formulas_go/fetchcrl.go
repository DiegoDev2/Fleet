package main

// FetchCrl - Retrieve certificate revocation lists (CRLs)
// Homepage: https://wiki.nikhef.nl/grid/FetchCRL3

import (
	"fmt"
	
	"os/exec"
)

func installFetchCrl() {
	// Método 1: Descargar y extraer .tar.gz
	fetchcrl_tar_url := "https://dist.eugridpma.info/distribution/util/fetch-crl3/fetch-crl-3.0.23.tar.gz"
	fetchcrl_cmd_tar := exec.Command("curl", "-L", fetchcrl_tar_url, "-o", "package.tar.gz")
	err := fetchcrl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fetchcrl_zip_url := "https://dist.eugridpma.info/distribution/util/fetch-crl3/fetch-crl-3.0.23.zip"
	fetchcrl_cmd_zip := exec.Command("curl", "-L", fetchcrl_zip_url, "-o", "package.zip")
	err = fetchcrl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fetchcrl_bin_url := "https://dist.eugridpma.info/distribution/util/fetch-crl3/fetch-crl-3.0.23.bin"
	fetchcrl_cmd_bin := exec.Command("curl", "-L", fetchcrl_bin_url, "-o", "binary.bin")
	err = fetchcrl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fetchcrl_src_url := "https://dist.eugridpma.info/distribution/util/fetch-crl3/fetch-crl-3.0.23.src.tar.gz"
	fetchcrl_cmd_src := exec.Command("curl", "-L", fetchcrl_src_url, "-o", "source.tar.gz")
	err = fetchcrl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fetchcrl_cmd_direct := exec.Command("./binary")
	err = fetchcrl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
