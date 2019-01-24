/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package models;

/**
 *
 * @author johancanales
 */
public class Boleto {
    private int idBoleto;
    private Usuario idUsuario;

    public Boleto(int idBoleto, Usuario idUsuario) {
        this.idBoleto = idBoleto;
        this.idUsuario = idUsuario;
    }

    public int getIdBoleto() {
        return idBoleto;
    }

    public Usuario getIdUsuario() {
        return idUsuario;
    }

    public void setIdBoleto(int idBoleto) {
        this.idBoleto = idBoleto;
    }

    public void setIdUsuario(Usuario idUsuario) {
        this.idUsuario = idUsuario;
    }
    
    
    
}
