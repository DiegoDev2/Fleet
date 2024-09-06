package main

// Heppdt2 - Translation for particle ID's to and from various MC generators and PDG standard
// Homepage: https://cdcvs.fnal.gov/redmine/projects/heppdt/wiki

import (
	"fmt"
	
	"os/exec"
)

func installHeppdt2() {
	// Método 1: Descargar y extraer .tar.gz
	heppdt2_tar_url := "https://cern.ch/lcgpackages/tarFiles/sources/HepPDT-2.06.01.tar.gz"
	heppdt2_cmd_tar := exec.Command("curl", "-L", heppdt2_tar_url, "-o", "package.tar.gz")
	err := heppdt2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	heppdt2_zip_url := "https://cern.ch/lcgpackages/tarFiles/sources/HepPDT-2.06.01.zip"
	heppdt2_cmd_zip := exec.Command("curl", "-L", heppdt2_zip_url, "-o", "package.zip")
	err = heppdt2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	heppdt2_bin_url := "https://cern.ch/lcgpackages/tarFiles/sources/HepPDT-2.06.01.bin"
	heppdt2_cmd_bin := exec.Command("curl", "-L", heppdt2_bin_url, "-o", "binary.bin")
	err = heppdt2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	heppdt2_src_url := "https://cern.ch/lcgpackages/tarFiles/sources/HepPDT-2.06.01.src.tar.gz"
	heppdt2_cmd_src := exec.Command("curl", "-L", heppdt2_src_url, "-o", "source.tar.gz")
	err = heppdt2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	heppdt2_cmd_direct := exec.Command("./binary")
	err = heppdt2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
