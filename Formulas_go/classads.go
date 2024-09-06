package main

// Classads - Classified Advertisements (used by HTCondor Central Manager)
// Homepage: https://research.cs.wisc.edu/htcondor/classad/

import (
	"fmt"
	
	"os/exec"
)

func installClassads() {
	// Método 1: Descargar y extraer .tar.gz
	classads_tar_url := "https://ftp.cs.wisc.edu/condor/classad/c++/classads-1.0.10.tar.gz"
	classads_cmd_tar := exec.Command("curl", "-L", classads_tar_url, "-o", "package.tar.gz")
	err := classads_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	classads_zip_url := "https://ftp.cs.wisc.edu/condor/classad/c++/classads-1.0.10.zip"
	classads_cmd_zip := exec.Command("curl", "-L", classads_zip_url, "-o", "package.zip")
	err = classads_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	classads_bin_url := "https://ftp.cs.wisc.edu/condor/classad/c++/classads-1.0.10.bin"
	classads_cmd_bin := exec.Command("curl", "-L", classads_bin_url, "-o", "binary.bin")
	err = classads_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	classads_src_url := "https://ftp.cs.wisc.edu/condor/classad/c++/classads-1.0.10.src.tar.gz"
	classads_cmd_src := exec.Command("curl", "-L", classads_src_url, "-o", "source.tar.gz")
	err = classads_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	classads_cmd_direct := exec.Command("./binary")
	err = classads_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
