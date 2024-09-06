package main

// Perl - Highly capable, feature-rich programming language
// Homepage: https://www.perl.org/

import (
	"fmt"
	
	"os/exec"
)

func installPerl() {
	// Método 1: Descargar y extraer .tar.gz
	perl_tar_url := "https://www.cpan.org/src/5.0/perl-5.38.2.tar.xz"
	perl_cmd_tar := exec.Command("curl", "-L", perl_tar_url, "-o", "package.tar.gz")
	err := perl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perl_zip_url := "https://www.cpan.org/src/5.0/perl-5.38.2.tar.xz"
	perl_cmd_zip := exec.Command("curl", "-L", perl_zip_url, "-o", "package.zip")
	err = perl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perl_bin_url := "https://www.cpan.org/src/5.0/perl-5.38.2.tar.xz"
	perl_cmd_bin := exec.Command("curl", "-L", perl_bin_url, "-o", "binary.bin")
	err = perl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perl_src_url := "https://www.cpan.org/src/5.0/perl-5.38.2.tar.xz"
	perl_cmd_src := exec.Command("curl", "-L", perl_src_url, "-o", "source.tar.gz")
	err = perl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perl_cmd_direct := exec.Command("./binary")
	err = perl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
}
