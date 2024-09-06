package main

// ErlangAT22 - Programming language for highly scalable real-time systems
// Homepage: https://www.erlang.org/

import (
	"fmt"
	
	"os/exec"
)

func installErlangAT22() {
	// Método 1: Descargar y extraer .tar.gz
	erlangat22_tar_url := "https://github.com/erlang/otp/releases/download/OTP-22.3.4.26/otp_src_22.3.4.26.tar.gz"
	erlangat22_cmd_tar := exec.Command("curl", "-L", erlangat22_tar_url, "-o", "package.tar.gz")
	err := erlangat22_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erlangat22_zip_url := "https://github.com/erlang/otp/releases/download/OTP-22.3.4.26/otp_src_22.3.4.26.zip"
	erlangat22_cmd_zip := exec.Command("curl", "-L", erlangat22_zip_url, "-o", "package.zip")
	err = erlangat22_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erlangat22_bin_url := "https://github.com/erlang/otp/releases/download/OTP-22.3.4.26/otp_src_22.3.4.26.bin"
	erlangat22_cmd_bin := exec.Command("curl", "-L", erlangat22_bin_url, "-o", "binary.bin")
	err = erlangat22_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erlangat22_src_url := "https://github.com/erlang/otp/releases/download/OTP-22.3.4.26/otp_src_22.3.4.26.src.tar.gz"
	erlangat22_cmd_src := exec.Command("curl", "-L", erlangat22_src_url, "-o", "source.tar.gz")
	err = erlangat22_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erlangat22_cmd_direct := exec.Command("./binary")
	err = erlangat22_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@1.1")
exec.Command("latte", "install", "openssl@1.1")
	fmt.Println("Instalando dependencia: wxwidgets")
exec.Command("latte", "install", "wxwidgets")
}
