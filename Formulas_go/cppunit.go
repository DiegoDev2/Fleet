package main

// Cppunit - Unit testing framework for C++
// Homepage: https://wiki.freedesktop.org/www/Software/cppunit/

import (
	"fmt"
	
	"os/exec"
)

func installCppunit() {
	// Método 1: Descargar y extraer .tar.gz
	cppunit_tar_url := "https://dev-www.libreoffice.org/src/cppunit-1.15.1.tar.gz"
	cppunit_cmd_tar := exec.Command("curl", "-L", cppunit_tar_url, "-o", "package.tar.gz")
	err := cppunit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppunit_zip_url := "https://dev-www.libreoffice.org/src/cppunit-1.15.1.zip"
	cppunit_cmd_zip := exec.Command("curl", "-L", cppunit_zip_url, "-o", "package.zip")
	err = cppunit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppunit_bin_url := "https://dev-www.libreoffice.org/src/cppunit-1.15.1.bin"
	cppunit_cmd_bin := exec.Command("curl", "-L", cppunit_bin_url, "-o", "binary.bin")
	err = cppunit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppunit_src_url := "https://dev-www.libreoffice.org/src/cppunit-1.15.1.src.tar.gz"
	cppunit_cmd_src := exec.Command("curl", "-L", cppunit_src_url, "-o", "source.tar.gz")
	err = cppunit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppunit_cmd_direct := exec.Command("./binary")
	err = cppunit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
