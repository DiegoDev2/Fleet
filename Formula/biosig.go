package main

// Biosig - Tools for biomedical signal processing and data conversion
// Homepage: https://biosig.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBiosig() {
	// Método 1: Descargar y extraer .tar.gz
	biosig_tar_url := "https://downloads.sourceforge.net/project/biosig/BioSig%20for%20C_C%2B%2B/src/biosig-2.6.1.src.tar.xz"
	biosig_cmd_tar := exec.Command("curl", "-L", biosig_tar_url, "-o", "package.tar.gz")
	err := biosig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	biosig_zip_url := "https://downloads.sourceforge.net/project/biosig/BioSig%20for%20C_C%2B%2B/src/biosig-2.6.1.src.tar.xz"
	biosig_cmd_zip := exec.Command("curl", "-L", biosig_zip_url, "-o", "package.zip")
	err = biosig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	biosig_bin_url := "https://downloads.sourceforge.net/project/biosig/BioSig%20for%20C_C%2B%2B/src/biosig-2.6.1.src.tar.xz"
	biosig_cmd_bin := exec.Command("curl", "-L", biosig_bin_url, "-o", "binary.bin")
	err = biosig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	biosig_src_url := "https://downloads.sourceforge.net/project/biosig/BioSig%20for%20C_C%2B%2B/src/biosig-2.6.1.src.tar.xz"
	biosig_cmd_src := exec.Command("curl", "-L", biosig_src_url, "-o", "source.tar.gz")
	err = biosig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	biosig_cmd_direct := exec.Command("./binary")
	err = biosig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: libb64")
	exec.Command("latte", "install", "libb64").Run()
	fmt.Println("Instalando dependencia: dcmtk")
	exec.Command("latte", "install", "dcmtk").Run()
	fmt.Println("Instalando dependencia: suite-sparse")
	exec.Command("latte", "install", "suite-sparse").Run()
}
