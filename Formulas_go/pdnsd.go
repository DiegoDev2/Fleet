package main

// Pdnsd - Proxy DNS server with permanent caching
// Homepage: https://web.archive.org/web/20201203080556/members.home.nl/p.a.rombouts/pdnsd/

import (
	"fmt"
	
	"os/exec"
)

func installPdnsd() {
	// Método 1: Descargar y extraer .tar.gz
	pdnsd_tar_url := "https://web.archive.org/web/20200323100335/members.home.nl/p.a.rombouts/pdnsd/releases/pdnsd-1.2.9a-par.tar.gz"
	pdnsd_cmd_tar := exec.Command("curl", "-L", pdnsd_tar_url, "-o", "package.tar.gz")
	err := pdnsd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdnsd_zip_url := "https://web.archive.org/web/20200323100335/members.home.nl/p.a.rombouts/pdnsd/releases/pdnsd-1.2.9a-par.zip"
	pdnsd_cmd_zip := exec.Command("curl", "-L", pdnsd_zip_url, "-o", "package.zip")
	err = pdnsd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdnsd_bin_url := "https://web.archive.org/web/20200323100335/members.home.nl/p.a.rombouts/pdnsd/releases/pdnsd-1.2.9a-par.bin"
	pdnsd_cmd_bin := exec.Command("curl", "-L", pdnsd_bin_url, "-o", "binary.bin")
	err = pdnsd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdnsd_src_url := "https://web.archive.org/web/20200323100335/members.home.nl/p.a.rombouts/pdnsd/releases/pdnsd-1.2.9a-par.src.tar.gz"
	pdnsd_cmd_src := exec.Command("curl", "-L", pdnsd_src_url, "-o", "source.tar.gz")
	err = pdnsd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdnsd_cmd_direct := exec.Command("./binary")
	err = pdnsd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
